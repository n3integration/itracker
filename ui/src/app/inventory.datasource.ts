import {DataSource} from '@angular/cdk/collections';
import {Inject, Injectable} from '@angular/core';

import {BehaviorSubject, Observable, of} from 'rxjs';
import {catchError} from 'rxjs/operators';

import {Item, ItemCollection} from './inventory.model';
import {InventoryService} from './services/inventory.service';

@Injectable({
    providedIn: 'root'
})
export class InventoryDataSource implements DataSource<ItemCollection> {

    private itemsSubject = new BehaviorSubject<ItemCollection[]>([]);

    constructor(@Inject(InventoryService) private inventoryService: InventoryService) {
    }

    connect(): Observable<ItemCollection[]> {
        return this.itemsSubject.asObservable();
    }

    disconnect(): void {
        this.itemsSubject.complete();
    }

    loadInventory() {
        this.inventoryService.getAll().pipe(
            catchError(() => of([]))
        ).subscribe(items => this.itemsSubject.next(items));
    }

    insert(item: Item) {
        this.itemsSubject.getValue().forEach((col) => {
            if (col.code === item.code) {
                col.addItem(item);
            }
        });
    }

    update(item: Item) {
        this.itemsSubject.getValue().forEach((col) => {
            if (col.code === item.code) {
                col.items.filter((i) => i.serial === item.serial)
                    .forEach((i) => {
                        i.facility = item.facility;
                        i.status = item.status;
                        i.submittedBy = item.submittedBy;
                    });
            }
        });
    }
}
