import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {NgModule} from '@angular/core';

import {MatTableModule} from '@angular/material/table';
import {MatListModule} from '@angular/material/list';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatButtonModule} from '@angular/material/button';
import {MatDialogModule} from '@angular/material/dialog';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatSelectModule} from '@angular/material/select';
import {MatSortModule} from '@angular/material/sort';
import {MatIconModule} from '@angular/material/icon';
import {MAT_SNACK_BAR_DEFAULT_OPTIONS, MatSnackBarModule} from '@angular/material/snack-bar';
import {MAT_RADIO_DEFAULT_OPTIONS, MatRadioModule} from '@angular/material/radio';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {InStockItemsListComponent} from './in-stock-items/in-stock-items-list.component';
import {AddItemDialogComponent} from './add-item/add-item-dialog.component';
import {InventoryService} from './services/inventory.service';
import {InventoryDataSource} from './inventory.datasource';
import {ItemCollectionListComponent} from './item-collection/item-collection-list.component';
import {UpdateItemDialogComponent} from './update-item/update-item-dialog.component';
import {ItemHistoryDialogComponent} from './item-history/item-history-dialog.component';

@NgModule({
    declarations: [
        AppComponent,
        AddItemDialogComponent,
        UpdateItemDialogComponent,
        ItemHistoryDialogComponent,
        InStockItemsListComponent,
        ItemCollectionListComponent,
    ],
    imports: [
        HttpClientModule,
        BrowserModule,
        AppRoutingModule,
        MatDialogModule,
        MatButtonModule,
        MatToolbarModule,
        MatTableModule,
        BrowserAnimationsModule,
        MatFormFieldModule,
        FormsModule,
        MatInputModule,
        MatSelectModule,
        ReactiveFormsModule,
        MatSortModule,
        MatSnackBarModule,
        MatListModule,
        MatIconModule,
        MatRadioModule,
    ],
    entryComponents: [
        AddItemDialogComponent,
        UpdateItemDialogComponent,
        ItemHistoryDialogComponent,
    ],
    providers: [
        InventoryService,
        InventoryDataSource,
        {provide: MAT_SNACK_BAR_DEFAULT_OPTIONS, useValue: {duration: 5000}},
        {provide: MAT_RADIO_DEFAULT_OPTIONS, useValue: {color: 'primary'}},
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
