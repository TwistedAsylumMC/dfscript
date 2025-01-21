declare namespace chat {
    interface Chat {
        write(bytes: Uint8Array): void;
        writeString(str: string): void;
        writet(translation: Translation, ...args: any[]): void;
        subscribe(subscriber: any): void; // TODO: Subscriber type
        subcribed(subscriber: any): boolean; // TODO: Subscriber type
        unsubscribe(subscriber: any): void; // TODO: Subscriber type
        close(): void;
    }
    const global: Chat;

    interface Translation {} // TODO
    function translate(): Translation;
}