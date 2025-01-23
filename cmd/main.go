package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
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
	createRuntime(srv, "examples", ".index.js")

	srv.Listen()
	for p := range srv.Accept() {
		h := runtime.PlayerJoin(p)
		p.Handle(h)
	}
}

type reloadCommand struct {
	srv *server.Server
}

func (r reloadCommand) Run(cmd.Source, *cmd.Output, *world.Tx) {
	if runtime != nil {
		runtime.Close()
	}
	createRuntime(r.srv, "examples", ".index.js")

	runtime.Loop().RunOnLoop(func(*goja.Runtime) {
		// Run inside the loop to ensure all onPlayerJoin handlers are registered.
		for player := range r.srv.Players(nil) {
			h := runtime.PlayerJoin(player)
			player.Handle(h)
		}
	})
}

func createRuntime(srv *server.Server, scriptDir, suffix string) {
	var err error
	runtime, err = dfscript.NewRuntime(srv)
	if err != nil {
		panic(fmt.Errorf("new runtime: %w", err))
	}
	wh := runtime.AddWorld(srv.World())
	srv.World().Handle(wh)

	err = filepath.Walk(scriptDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), suffix) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		slog.Debug("Running script", "path", path)
		if ok := runtime.Run(string(data)); !ok {
			return fmt.Errorf("failed to run %s", path)
		}
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("walk %s: %w", scriptDir, err))
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
