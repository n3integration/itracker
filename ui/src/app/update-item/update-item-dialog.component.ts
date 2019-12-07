import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Observable} from 'rxjs';

import {Item} from '../inventory.model';
import {InventoryService} from '../services/inventory.service';
import {InventoryDataSource} from '../inventory.datasource';

interface UpdateTO {
    item: Item;
    op: string;
    disabled: boolean;
}

@Component({
    selector: 'app-update-item-dialog',
    templateUrl: 'update-item-dialog.component.html',
})
export class UpdateItemDialogComponent {
    constructor(
        private snackBar: MatSnackBar,
        private dialogRef: MatDialogRef<UpdateItemDialogComponent>,
        @Inject(InventoryService) private service: InventoryService,
        @Inject(InventoryDataSource) private datasource: InventoryDataSource,
        @Inject(MAT_DIALOG_DATA) public data: UpdateTO) {
    }

    submit(data: UpdateTO): void {
        if (data.disabled) {
            return;
        }

        this.dialogRef.close();
        let itemObservable: Observable<Item>;

        switch (data.op) {
            case 'update':
                itemObservable = this.service.updateStatus(data.item);
                break;
            case 'transfer':
                itemObservable = this.service.transfer(data.item);
                break;
        }

        itemObservable.subscribe(
            (result) => {
                this.datasource.update(result);
                this.snackBar.open(`${data.op} complete`);
            },
            (err) => this.snackBar.open(err.error.message),
        );
    }

    onNoClick() {
        this.dialogRef.close();
    }
}
