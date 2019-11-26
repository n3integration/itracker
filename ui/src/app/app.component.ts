import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from '@angular/material/dialog';

import {InventoryService} from './services/inventory.service';
import {Item} from './item.model';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.sass']
})
export class AppComponent {

    constructor(public dialog: MatDialog) {
    }

    openDialog() {
        const dialogRef = this.dialog.open(AddItemDialogComponent, {
            width: '400px',
            data: {facility: 'Warehouse'},
        });

        dialogRef.afterClosed().subscribe(result => {

        });
    }
}

export interface Template {
    name: string;
    codes: Code[];
}

export interface Code {
    name: string;
    category: string;
}

@Component({
    selector: 'app-add-item-dialog',
    templateUrl: 'add-item-dialog.html',
})
export class AddItemDialogComponent {
    manufacturers: Template[] = [
        // tslint:disable:max-line-length
        {name: 'SPOD', codes: [{name: 'POD300-ARB', category: 'Air Accessory Kit'}]},
        {
            name: 'VIAIR', codes: [{name: 'V/A00025', category: 'Air Accessory Kit'},
                {name: 'V/A00027', category: 'Air Accessory Kit'},
                {name: 'V/A00029', category: 'Air Accessory Kit'},
                {name: 'V/A20052', category: 'Air Accessory Kit'},
                {name: 'V/A20053', category: 'Air Accessory Kit'},
                {name: 'V/A20055', category: 'Air Accessory Kit'},
                {name: 'V/A90007', category: 'Air Accessory Kit'},
                {name: 'V/A92621', category: 'Air Accessory Kit'},
                {name: 'V/A92622', category: 'Air Accessory Kit'},
                {name: 'V/A92623', category: 'Air Accessory Kit'},
                {name: 'V/A92625', category: 'Air Accessory Kit'},
                {name: 'V/A92626', category: 'Air Accessory Kit'},
                {name: 'V/A92627', category: 'Air Accessory Kit'},
                {name: 'V/A92630', category: 'Air Accessory Kit'},
                {name: 'V/A92631', category: 'Air Accessory Kit'},
                {name: 'V/A92635', category: 'Air Accessory Kit'},
                {name: 'V/A92595', category: 'Air Accessory Kit'}]
        },
        {name: 'Smittybilt', codes: [{name: 'S/B2781BAG', category: 'Air Compressor Carry Bag'}]},
        {name: 'TeraFlex', codes: [{name: 'TER1184120', category: 'Air Compressor Mounting Bracket'}]},
    ];

    constructor(
        public dialogRef: MatDialogRef<AddItemDialogComponent>,
        @Inject(MAT_DIALOG_DATA) public data: Item,
        @Inject(InventoryService) public service: InventoryService) {
    }

    getCodes(name): Code[] {
        const manufacturer = this.manufacturers.find(v => v.name === name);
        if (manufacturer) {
            return manufacturer.codes;
        }
        return [];
    }

    updateCode(data: Item) {
        const codes = this.getCodes(data.manufacturer);
        if (codes.length) {
            data.code = codes[0].name;
            this.updateCategory(data);
        }
    }

    updateCategory(data: Item) {
        const codes = this.getCodes(data.manufacturer);
        if (codes.length) {
            const selected = codes.find(v => v.name === data.code);
            if (selected) {
                data.category = selected.category;
            }
        }
    }

    submit(item: Item): void {
        this.dialogRef.close();
        this.service.addItem(item);
    }

    onNoClick(): void {
        this.dialogRef.close();
    }
}
