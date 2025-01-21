declare namespace title {
    interface Title {
        text(): string;
        subtitle(): string;
        withSubtitle(...args: any[]): Title;
        actionText(): string;
        withActionText(...args: any[]): Title;
        duration(): Date;
        withDuration(duration: Date): Title;
        fadeInDuration(): Date;
        withFadeInDuration(duration: Date): Title;
        fadeOutDuration(): Date;
        withFadeOutDuration(duration: Date): Title;
    }
    function create(...args: any[]): Title;
}