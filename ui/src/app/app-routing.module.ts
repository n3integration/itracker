import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {InStockItemsListComponent} from './in-stock-items/in-stock-items-list.component';
import {ItemCollectionListComponent} from './item-collection/item-collection-list.component';

const routes: Routes = [
  {path: '', redirectTo: 'inStock', pathMatch: 'full'},
  {path: 'inStock', component: InStockItemsListComponent, data: {animation: 'Home'}},
  {path: 'items/:code', component: ItemCollectionListComponent, data: {animation: 'Item'}},
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
