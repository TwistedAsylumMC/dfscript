/// <reference path="player.d.ts" />
/// <reference path="world.d.ts" />

declare namespace form {
    interface Form {}

    interface CustomForm extends Form {
        label(text: string): CustomForm;
        input(text: string, defaultText: string, placeholder: string): CustomForm;
        toggle(text: string, defaultValue: boolean): CustomForm;
        slider(text: string, min: number, max: number, step: number, defaultValue: number): CustomForm;
        stepSlider(text: string, steps: string[], defaultIndex: number): CustomForm;
        dropdown(text: string, options: string[], defaultIndex: number): CustomForm;
        onSubmit(callback: (tx: world.Tx, player: player.Player, data: any[], closed: boolean) => void): CustomForm;
    }
    function custom(title: string): CustomForm;

    interface MenuForm extends Form {
        button(text: string, image?: string, submit?: (tx: world.Tx, player: player.Player) => void): MenuForm;
        onSubmit(callback: (tx: world.Tx, player: player.Player, index: number, closed: boolean) => void): MenuForm;
    }
    function menu(title: string, body?: string): MenuForm;

    interface ModalForm extends Form {
        yes(text: string, submit?: (tx: world.Tx, player: player.Player) => void): ModalForm;
        no(text: string, submit?: (tx: world.Tx, player: player.Player) => void): ModalForm;
        onSubmit(callback: (tx: world.Tx, player: player.Player, index: number, closed: boolean) => void): ModalForm;
    }
    function modal(text: string, body?: string): ModalForm;
}