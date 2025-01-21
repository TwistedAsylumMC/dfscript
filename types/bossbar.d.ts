declare namespace bossbar {
    interface Colour {}
    const colour: {
        blue: Colour;
        green: Colour;
        grey: Colour;
        purple: Colour;
        red: Colour;
        white: Colour;
        yellow: Colour;
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