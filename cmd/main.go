package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/pelletier/go-toml"
	"github.com/twistedasylummc/dfscript"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var runtime *dfscript.Runtime

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	chat.Global.Subscribe(chat.StdoutSubscriber{})
	conf, err := readConfig(slog.Default())
	if err != nil {
		panic(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	cmd.Register(cmd.New("reload", "Reloads all scripts", nil, reloadCommand{srv: srv}))

	runtime, err = dfscript.NewRuntime(srv)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk("examples", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), ".index.js") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		fmt.Println("Running", path)
		if ok := runtime.Run(string(data)); !ok {
			return fmt.Errorf("failed to run %s", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	srv.Listen()
	for p := range srv.Accept() {
		h := runtime.PlayerJoin(p)
		p.Handle(h)
	}
}

type reloadCommand struct {
	srv *server.Server
}

func (r reloadCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	//if runtime != nil {
	//	runtime.Close()
	//}

	var err error
	runtime, err = dfscript.NewRuntime(r.srv)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk("examples", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), ".index.js") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		fmt.Println("Running", path)
		if ok := runtime.Run(string(data)); !ok {
			return fmt.Errorf("failed to run %s", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for player := range r.srv.Players(tx) {
		h := runtime.PlayerJoin(player)
		player.Handle(h)
	}
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log *slog.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
