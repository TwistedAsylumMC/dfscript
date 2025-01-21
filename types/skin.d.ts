declare namespace skin {
    interface AnimationType {}
    const animationType: {
        head: AnimationType;
        body32x32: AnimationType;
        body128x128: AnimationType;
    }
    interface Animation {
        pix: Uint8Array;
        frameCount: number;
        animationExpression: number;

        type(): AnimationType;
        // TODO: colorModel
        // TODO: bounds
        // TODO: at
    }
    function animation(width: number, height: number, expression: number, type: AnimationType): Animation;

    interface Cape {
        pix: Uint8Array;
        // TODO: colorModel
        // TODO: bounds
        // TODO: at
    }
    function cape(width: number, height: number): Cape;

    interface ModelConfig {
        default: string;
        animatedFace: string;

        encode(): Uint8Array;
    }
    function modelConfig(def: string, face: string): ModelConfig;
    function decodeModelConfig(data: Uint8Array): ModelConfig;

    interface Skin {
        pix: Uint8Array;
        persona: boolean;
        playFabID: string;
        modelConfig: ModelConfig;
        model: Uint8Array;
        cape: Cape;
        animations: Animation[];

        // TODO: colorModel
        // TODO: bounds
        // TODO: at
    }
    function create(width: number, height: number): Skin;
}