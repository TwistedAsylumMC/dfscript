console.log(biome.badlands.string())
console.log(cube.rotation(60, 90).elem())
console.log(block.air().encodeBlock())
console.log(block.anvil({type: block.AnvilType.SlightlyDamaged}).facing)

server.onPlayerJoin(player => {
    player.message(text.emerald + `Welcome, ${player.name}`)
    player.sendTitle(title.create(text.diamond + "Welcome!").withSubtitle(text.gold + "sent by javascript"))
    player.showCoordinates()
    player.setGameMode(world.GameMode.Survival)
    player.inventory().addItem(item.stack(item.apple(), 8))
    setTimeout(() => {
        player.message(text.colourf("<red>%s</red>", "Delayed message"))
        let f = form.menu("title")
            .button("test")
            .button("test2", "textures/items/diamond", (player) => {
                player.message("You pressed test2")
                player.inventory().addItem(item.stack(item.splashPotion({type: potion.wither}), 1))
            })
            .button("place blocks", "", (player) => {
                let world = player.world();
                world.setBlock(cube.pos(0, -60, 0), block.stone())
                world.setBlock(cube.pos(1, -60, 0), block.stone())
                world.setBlock(cube.pos(2, -60, 0), block.stone())
                world.setBlock(cube.pos(3, -60, 0), block.stone())
                world.setBlock(cube.pos(4, -60, 0), block.stone())
                world.setBlock(cube.pos(5, -60, 0), block.stone())
            })
            .onSubmit((player, index, closed) => {
                player.message(`You pressed index ${index}. Closed: ${closed}`)
            })
        player.sendForm(f)
    }, 1000)

    player.onHeldSlotChange(e => {
        player.message("Changed slot from", e.from, "to", e.to)
        if (e.to === 4) {
            e.cancel()
        }
    })

    let firstMessage = true;
    player.onChat(e => {
        if (firstMessage) {
            firstMessage = false;
            return;
        }
        e.cancel()
        chat.global.writeString(`[${player.name}] -> ${e.message}`)
    })
})

// setInterval(() => {
//     server.players().forEach(player => {
//         player.message("Hello, everyone!")
//     })
// }, 2000)