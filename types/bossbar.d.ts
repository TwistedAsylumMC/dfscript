declare namespace bossbar {
    class Colour {
        static readonly blue: Colour;
        static readonly green: Colour;
        static readonly grey: Colour;
        static readonly purple: Colour;
        static readonly red: Colour;
        static readonly white: Colour;
        static readonly yellow: Colour;

        uint8(): number;
    }

    interface BossBar {
        text(): string;
        healthPercentage(): number;
        withHealthPercentage(percentage: number): BossBar;
        colour(): Colour;
        withColour(colour: Colour): BossBar;
    }
    function create(...args: any[]): BossBar;
}