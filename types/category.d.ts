declare namespace category {
    interface Category {
        uint8(): number;
        withGroup(group: string): Category;
        string(): string;
        group(): string;
    }

    const construction: Category;
    const nature: Category;
    const equipment: Category;
    const items: Category;
}