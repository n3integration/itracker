import {Component, OnInit} from '@angular/core';
import {InventoryService} from './services/inventory.service';
import {ItemCollection} from './item.model';

@Component({
    selector: 'app-item-list',
    templateUrl: './item-list.component.html',
    providers: [InventoryService],
})
export class ItemListComponent implements OnInit {
    items: ItemCollection[];
    columns: string[] = ['code', 'item', 'category', 'manufacturer', 'quantity'];

    constructor(private service: InventoryService) {
    }

    ngOnInit(): void {
        this.items = this.service.getItems();
    }
}
