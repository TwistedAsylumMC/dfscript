console.log(biome.badlands.string())
console.log(cube.rotation(60, 90).elem())
console.log(block.air().encodeBlock())
console.log(block.anvil({type: block.AnvilType.SlightlyDamaged}).facing)

let f = form.menu("title")
    .button("test")
    .button("test2", "textures/items/diamond", (tx, player) => {
        player.message("You pressed test2")
        player.inventory().addItem(item.stack(item.splashPotion({type: potion.wither}), 1))
    })
    .button("place blocks", "", (tx, player) => {
        for (let i = 0; i < 10; i++) {
            tx.setBlock(cube.posFromVec3(player.position()).add(cube.pos(i, 0, 0)), block.stone())
        }
    })
    .onSubmit((tx, player, index, closed) => {
        player.message(`You pressed index ${index}. Closed: ${closed}`)
    })

event.onPlayerJoin(player => {
    console.log(player.xuid())
    player.message(text.emerald + `Welcome, ${player.name()}`)
    player.sendTitle(title.create(text.diamond + "Welcome!").withSubtitle(text.gold + "sent by javascript"))
    player.showCoordinates()
    player.setGameMode(world.GameMode.Survival)
    player.inventory().addItem(item.stack(item.apple(), 8))
    setTimeout(() => {
        player.h().execWorld((tx, player) => {
            player.message(text.colourf("<red>%s</red>", "Delayed message"))
            player.sendForm(f)
        })
    }, 1000)
})

event.onPlayerHeldSlotChange(e => {
    e.player.message("Changed slot from", e.from, "to", e.to)
    if (e.to === 4) {
        e.cancel()
    }
})

let firstMessage = true;
event.onPlayerChat(e => {
    if (firstMessage) {
        firstMessage = false;
        return;
    }
    e.cancel()
    chat.global.writeString(`[${e.player.name()}] -> ${e.message}`)
})

setInterval(() => {
    server.players()(player => {
        player.message("Hello, everyone!")
    })
}, 2000)