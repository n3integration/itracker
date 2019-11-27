import {Component, Inject, OnInit, ViewChild} from '@angular/core';
import {Router} from '@angular/router';

import {MatSort} from '@angular/material/sort';
import {MatDialog} from '@angular/material/dialog';
import {MatTableDataSource} from '@angular/material/table';

import {InventoryService} from '../services/inventory.service';
import {AddItemDialogComponent} from '../add-item/add-item-dialog.component';
import {InventoryDataSource} from '../inventory.datasource';
import {ItemCollection} from '../inventory.model';

@Component({
    selector: 'app-item-list',
    templateUrl: './in-stock-items-list.component.html',
})
export class InStockItemsListComponent implements OnInit {
    private columns: string[] = ['code', 'item', 'category', 'manufacturer', 'quantity'];
    private dataSource = new MatTableDataSource<ItemCollection>([]);

    @ViewChild(MatSort, {static: true}) sort: MatSort;

    constructor(@Inject(InventoryService) private inventoryService: InventoryService,
                @Inject(InventoryDataSource) private ds: InventoryDataSource,
                private dialog: MatDialog,
                private router: Router) {
    }

    ngOnInit(): void {
        this.dataSource.sort = this.sort;
        this.ds.loadInventory();
        this.ds.connect().subscribe((data) => this.dataSource.data = data);
    }

    goto(code: string): void {
        this.router.navigate(['/items', encodeURIComponent(code)])
            .then((result) => result ? null : console.log(`failed to navigate to items/${code}`));
    }

    openDialog() {
        const dialogRef = this.dialog.open(AddItemDialogComponent, {
            width: '400px',
            data: {facility: 'Warehouse'},
        });
    }
}
