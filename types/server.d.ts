/// <reference path="player.d.ts" />

declare namespace server {
    function playerCount(): number;
    function maxPlayerCount(): number;
    function players(): player.Player[];
    function player(uuid: string): player.Player | null;
    function playerByName(name: string): player.Player | null;
    function playerByXUID(xuid: string): player.Player | null;
    function onPlayerJoin(callback: (player: player.Player) => void): void;
    function world(name: string): world.World | null;
    function worlds(): world.World[];
}
