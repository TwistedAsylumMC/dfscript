/// <reference path="entity.d.ts" />
/// <reference path="player.d.ts" />
/// <reference path="uuid.d.ts" />
/// <reference path="world.d.ts" />

declare namespace server {
    function playerCount(): number;
    function maxPlayerCount(): number;
    function players(tx?: world.Tx): (player: player.Player) => boolean;
    function player(uuid: uuid.UUID): entity.Handle<player.Player> | null;
    function playerByName(name: string): entity.Handle<player.Player> | null;
    function playerByXUID(xuid: string): entity.Handle<player.Player> | null;
    function world(name: string): world.World | null;
    function worlds(): world.World[];
}
