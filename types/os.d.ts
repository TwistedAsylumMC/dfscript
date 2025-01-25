declare namespace os {
    function readFile(path: string): string;
    function writeFile(path: string, data: string): void;
}