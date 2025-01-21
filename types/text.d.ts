declare namespace text {
    const black: string;
    const darkBlue: string;
    const darkGreen: string;
    const darkAqua: string;
    const darkRed: string;
    const darkPurple: string;
    const orange: string;
    const grey: string;
    const darkGrey: string;
    const blue: string;
    const green: string;
    const aqua: string;
    const red: string;
    const purple: string;
    const yellow: string;
    const white: string;
    const darkYellow: string;
    const quartz: string;
    const iron: string;
    const netherite: string;
    const obfuscated: string;
    const bold: string;
    const redstone: string;
    const copper: string;
    const italic: string;
    const gold: string;
    const emerald: string;
    const reset: string;
    const diamond: string;
    const lapis: string;
    const amethyst: string;
    const resin: string;

    function ansi(...args: any[]): string;
    function clean(string: string): string;
    function colourf(format: string, ...args: any[]): string;
}