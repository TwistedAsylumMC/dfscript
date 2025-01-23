/// <reference path="entity.d.ts" />
/// <reference path="player.d.ts" />
/// <reference path="world.d.ts" />

declare namespace server {
    function playerCount(): number;
    function maxPlayerCount(): number;
    function players(): (player: player.Player) => boolean;
    function player(uuid: string): entity.Handle<player.Player> | null;
    function playerByName(name: string): entity.Handle<player.Player> | null;
    function playerByXUID(xuid: string): entity.Handle<player.Player> | null;
    function world(name: string): world.World | null;
    function worlds(): world.World[];
}
