const players = {}

const clearPlayer = player => {
    const playerData = players[player.name()];
    if (playerData) {
        for (const pos of playerData.blocks) {
            let b = playerData.replaced[pos] ? playerData.replaced[pos] : block.air()
            player.tx().setBlock(pos, b)
        }
        playerData.blocks = []
        playerData.replaced = {}
    }
}

function canReplace(b) {
    const [name] = b.encodeBlock()
    return [
        "minecraft:air", "minecraft:water", "minecraft:lava", "minecraft:deadbush", "minecraft:dandelion", "minecraft:poppy",
        "minecraft:blue_orchid", "minecraft:allium", "minecraft:azure_bluet", "minecraft:red_tulip", "minecraft:orange_tulip",
        "minecraft:white_tulip", "minecraft:pink_tulip", "minecraft:oxeye_daisy", "minecraft:cornflower", "minecraft:lily_of_the_valley",
        "minecraft:wither_rose", "minecraft:sunflower", "minecraft:lilac", "minecraft:rose_bush", "minecraft:peony",
        "minecraft:tall_grass", "minecraft:large_fern", "minecraft:fire", "minecraft:soul_fire", "minecraft:flowing_water",
        "minecraft:flowing_lava"
    ].includes(name)
}

event.onPlayerQuit(e => {
    clearPlayer(e.player)
    delete players[e.player.name()]
})

event.onPlayerChat(e => {
    if (e.message === ".magiccarpet" || e.message === ".mc") {
        e.cancel()
        if (players[e.player.name()]) {
            clearPlayer(e.player)
            delete players[e.player.name()]
            e.player.message(text.red + "Your magic carpet has been removed.")
            return
        }
        players[e.player.name()] = {radius: 2, lastPos: null, blocks: [], replaced: {}}
        e.player.message(text.green + "Your magic carpet has been created.")
    }
})

event.onPlayerMove(e => {
    if (!players[e.player.name()]) return
    const playerData = players[e.player.name()]
    let pos = cube.posFromVec3(e.player.position()).side(cube.Face.Down)
    if (e.player.rotation().pitch() > 75) {
       pos = pos.side(cube.Face.Down)
    }
    if (playerData.lastPos === pos) return

    clearPlayer(e.player)
    e.player.resetFallDistance();
    for (let x = -playerData.radius; x <= playerData.radius; x++) {
        for (let z = -playerData.radius; z <= playerData.radius; z++) {
            const blockPos = pos.add(cube.pos(x, 0, z))
            const b = e.tx.block(blockPos)
            if (!canReplace(b)) continue
            playerData.blocks.push(blockPos)
            playerData.replaced[blockPos] = b
            e.tx.setBlock(blockPos, block.glass())
        }
    }
})
