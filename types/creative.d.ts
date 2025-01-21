/// <reference path="item.d.ts" />

declare namespace creative {
    function items(): item.Stack[];
    function registerItem(item: item.Stack): void;
}