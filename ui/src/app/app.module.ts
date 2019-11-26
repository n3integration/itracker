import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {NgModule} from '@angular/core';

import {MatTableModule} from '@angular/material/table';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatButtonModule} from '@angular/material/button';
import {MatDialogModule} from '@angular/material/dialog';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatSelectModule} from '@angular/material/select';

import {AppRoutingModule} from './app-routing.module';
import {AddItemDialogComponent, AppComponent} from './app.component';
import {ItemListComponent} from './item-list.component';
import {InventoryService} from './services/inventory.service';


@NgModule({
    declarations: [
        AppComponent,
        AddItemDialogComponent,
        ItemListComponent
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
    ],
    entryComponents: [
        AddItemDialogComponent,
    ],
    providers: [
        InventoryService,
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
