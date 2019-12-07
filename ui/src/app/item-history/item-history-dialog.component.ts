import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {MatSnackBar} from '@angular/material/snack-bar';

import {InventoryService} from '../services/inventory.service';
import {Item, ItemHistory} from '../inventory.model';

@Component({
    selector: 'app-item-history-dialog',
    templateUrl: 'item-history-dialog.component.html',
})
export class ItemHistoryDialogComponent implements OnInit {

    history: ItemHistory[];

    constructor(
        private snackBar: MatSnackBar,
        private dialogRef: MatDialogRef<ItemHistoryDialogComponent>,
        @Inject(InventoryService) private service: InventoryService,
        @Inject(MAT_DIALOG_DATA) public data) {
    }

    ngOnInit(): void {
        this.history = [];
        this.service.getHistory(this.data.item).subscribe(
            (history: ItemHistory[]) => this.history.push(...history),
            (err) => this.snackBar.open(err.error.message));
    }

    getMessage(itemHistory: ItemHistory): string {
        const index = this.history.findIndex((h) => itemHistory.txId === h.txId);
        if (index === 0) {
            return 'Item was first tracked';
        }

        const item: Item = JSON.parse(itemHistory.value);
        if (item.status !== 1) {
            return 'Item was sold';
        }
        return `Item was transferred to ${item.facility}`;
    }

    onNoClick() {
        this.dialogRef.close();
    }
}
