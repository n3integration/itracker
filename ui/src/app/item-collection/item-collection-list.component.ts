import {Component, Inject, OnInit, ViewChild} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';

import {MatTableDataSource} from '@angular/material/table';
import {MatSort} from '@angular/material/sort';
import {MatDialog} from '@angular/material/dialog';

import {InventoryService} from '../services/inventory.service';
import {Item, ItemCollection} from '../inventory.model';
import {InventoryDataSource} from '../inventory.datasource';
import {UpdateItemDialogComponent} from '../update-item/update-item-dialog.component';

@Component({
    selector: 'app-item-collection-list',
    templateUrl: './item-collection-list.component.html',
})
export class ItemCollectionListComponent implements OnInit {
    private columns: string[] = ['serial', 'facility', 'status', 'submittedBy'];

    private collection$: ItemCollection;
    private dataSource: MatTableDataSource<Item>;

    @ViewChild(MatSort, {static: true}) sort: MatSort;

    constructor(private router: Router,
                private route: ActivatedRoute,
                private dialog: MatDialog,
                @Inject(InventoryService) private service: InventoryService,
                @Inject(InventoryDataSource) private ds: InventoryDataSource) {
        this.dataSource = new MatTableDataSource<Item>([]);
    }

    ngOnInit(): void {
        this.dataSource.sort = this.sort;
        const code = this.route.snapshot.paramMap.get('code');
        this.service.get(decodeURIComponent(code))
            .subscribe((col) => {
                this.collection$ = col;
                this.dataSource.data = this.collection$.items;
            });
    }

    goBack() {
        history.back();
    }

    openDialog(item: Item) {
        const dialogRef = this.dialog.open(UpdateItemDialogComponent, {
            width: '420px',
            data: {item: Object.assign({}, item), op: 'update', disabled: item.status === 0},
        });
    }
}
