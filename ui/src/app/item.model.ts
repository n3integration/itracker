export class Item {
    serial: string;
    manufacturer: string;
    code: string;
    category: string;
    facility: string;
    status: number;
    submittedBy: string;
}

export class ItemHistory {
    txId: string;
    value: JSON;
    timestamp: string;
}

export class ItemCollection {
    code: string;
    title: string;
    category: string;
    manufacturer: string;
    overview: string;
    imageUrl: string;
    items: Item[];

    constructor(code: string, title: string, imageUrl: string, overview: string) {
        this.code = code;
        this.title = title;
        this.overview = overview;
        this.imageUrl = imageUrl;
        this.items = [];
    }

    inStock(): number {
        const reducer = (accum, i) => i.status === 1 ? 1 + accum : 0;
        return this.items.reduce(reducer, 0);
    }

    total(): number {
        return this.items.length;
    }
}
