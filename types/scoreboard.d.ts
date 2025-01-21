declare namespace scoreboard {
    interface Scoreboard {
        name(): string;
        write(bytes: Uint8Array): void;
        writeString(string: string): void;
        set(index: number, string: string): void;
        remove(index: number): void;
        removePadding(): void;
        lines(): string[];
    }
    function create(...args: any[]): Scoreboard;
}