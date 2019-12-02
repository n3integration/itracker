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
    value: string;
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

    addItem(item: Item) {
        this.items.push(item);
    }

    inStock(): number {
        let sum = 0;
        this.items.forEach((i) => {
            if (i.status === 1) {
                sum += 1;
            }
        });
        return sum;
    }

    total(): number {
        return this.items.length;
    }
}
