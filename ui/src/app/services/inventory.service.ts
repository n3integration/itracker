import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

import {Observable, of} from 'rxjs';
import {map} from 'rxjs/operators';

import {Item, ItemCollection, ItemHistory} from '../inventory.model';

interface UpdateRequest {
    op: string;
    value?: string | null;
}

@Injectable({
    providedIn: 'root'
})
export class InventoryService {

    private baseUrl = 'http://localhost:8020';
    private items: ItemCollection[] = [
        // tslint:disable:max-line-length
        new ItemCollection('POD300-ARB', 'Adapter Harness for ARB Compressor - 300-ARB', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//78/16/77/b7/781677b73eaffdc105e8a25f8722d0def3712177.jpg', ''),
        new ItemCollection('V/A00025', '5-in-1 Inline Inflation/Deflation Coil Hose - 25', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//2f/65/b2/1d/2f65b21ddd185947cc40f4c2d2f90ebdeb7564d9.jpg', 'Open-ended air lines are meant for use with "H" model compressors, as well as portable compressors. They are not to be used with air tanks since they will allow air to flow freely from the air chuck as needed to fill tires and other inflatables...'),
        new ItemCollection('V/A00029', '5-in-1 Inline Inflation/Deflation Coil Hose - 29', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//8a/35/5f/69/8a355f69ee7f129d7239416c4b654e0ff55f7f2e.jpg', 'Features: Continuous Deflation-Push down and turn/lock the deflator collar. Intermittent Deflation-Push the deflator collar down to deflate, release to stop deflation. 25-ft. extension Coil Hose-Take the hose to your work area easily...'),
        new ItemCollection('V/A20055', 'Onboard Air Hookup Kit - 20055', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//9b/3d/2f/91/9b3d2f91d56683c8655187a7bcc4783148554a24.jpg', 'Custom designed to work with VIAIR Compressors. They each come with a Pressure Switch with Built-in Relay, an illuminated dash panel gauge with built-in ON/OFF switch, 20 ft. of positive extension wire, 20 ft. of air delivery line tubing...'),
        new ItemCollection('V/A90007', 'Air Source Relocation Kit - 90007', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//7b/aa/27/c4/7baa27c45822691e87ca90dd56c812037650e8da.jpg', 'Place your air source just about anywhere you want it with this relocation kit from VIAIR'),
        new ItemCollection('V/A20052', 'Onboard Air Hookup Kit - 20052', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//d9/1e/55/db/d91e55dbb3b006c14a2aac3609e2bab78a081021.jpg', 'Specially designed to work with VIAIR Air Compressors, these kits include a pressure switch, illuminated dash panel gauge with on/off switch, 20 ft. of air line, and all necessary fittings...'),
        new ItemCollection('V/A00027', '5-in-1 Inline Inflation/Deflation Coil Hose - 27', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//8a/35/5f/69/8a355f69ee7f129d7239416c4b654e0ff55f7f2e.jpg', 'Features: Continuous Deflation-Push down and turn/lock the deflator collar. Intermittent Deflation-Push the deflator collar down to deflate, release to stop deflation. 25-ft. extension Coil Hose-Take the hose to your work area easily...'),
        new ItemCollection('V/A20053', 'Onboard Air Hookup Kit - 20053', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//9b/3d/2f/91/9b3d2f91d56683c8655187a7bcc4783148554a24.jpg', 'Onboard Air Hookup Kits - are custom designed to work with VIAIR Compressors. They each come with a Pressure Switch with Built-in Relay, an illuminated dash panel gauge with built-in ON/OFF switch, 20 ft. of positive extension wire, 20 ft. of air delivery...'),
        new ItemCollection('S/B2781BAG', 'Compressor Storage Bag - 2781BAG', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//b4/78/47/39/b4784739c92dd7044987958037f47e6eba419649.jpg', 'Smittybilt bag for the S/B2781 moblie air compressor'),
        new ItemCollection('V/A92623', 'Direct Inlet Air Filter Assembly - 92623', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//2a/e6/5c/8d/2ae65c8d57f0a9c56ed468730a632d2483213411.jpg', 'Replacement direct-mount air filter assembly (on compressor) with 1/4" Male NPT'),
        new ItemCollection('V/A92630', 'Metal Direct Inlet Air Filter Assembly - 92630', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//41/fc/d9/70/41fcd970e013301530d63275f11953b45650dd2e.jpg', 'Replacement direct-mount air filter assembly with 1/4" NPT (metal housing)'),
        new ItemCollection('V/A92626', 'Dual Stage Air Filter Element - 92626', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//38/38/69/04/383869040fa79b8b4933279f2d6703c560302ee2.jpg', 'Air filter elements should be replaced periodically depending on frequency of use and operating environment. For use with plastic air filter housings.'),
        new ItemCollection('V/A92635', 'Direct Inlet Air Filter Assembly - 92635', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//41/fc/d9/70/41fcd970e013301530d63275f11953b45650dd2e.jpg', 'Replacement direct-mount air filter assembly with 1/2" NPT (metal housing)'),
        new ItemCollection('V/A92622', 'Remote Intake Air Filter Assembly - 92622', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//9f/31/60/c5/9f3160c504d8860c9d46328c0d89ac03cc25a21e.jpg', 'Remote Intake Air Filter Assembly, Plastic Housing (1/4" x 3/8" Tube Fitting, NPT)'),
        new ItemCollection('V/A92625', 'Remote Intake Air Filter Assembly - 92625', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//9f/31/60/c5/9f3160c504d8860c9d46328c0d89ac03cc25a21e.jpg', 'Remote Intake Air Filter Assembly, Plastic Housing (3/8" x 1/2" Tube Fitting, NPT)'),
        new ItemCollection('V/A92627', 'Direct Inlet Air Filter Assembly - 92627', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//41/fc/d9/70/41fcd970e013301530d63275f11953b45650dd2e.jpg', 'Replacement direct-mount air filter assembly with 3/8" NPT thread (metal housing)'),
        new ItemCollection('V/A92595', 'Dual Stage Air Filter Elements - 92595', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//38/38/69/04/383869040fa79b8b4933279f2d6703c560302ee2.jpg', 'Dual Stage Air Filter Elements'),
        new ItemCollection('V/A92631', 'Metal Remote Inlet Air Filter Assembly - 92631', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//dd/86/ec/d7/dd86ecd7871255f6d5bc8ec9df09a9985c13c542.jpg', 'Replacement remote-mount air filter assembly with 1/4" NPT (metal housing)'),
        new ItemCollection('V/A92621', 'Remote Intake Air Filter Assembly - 92621', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//9f/31/60/c5/9f3160c504d8860c9d46328c0d89ac03cc25a21e.jpg', 'Remote Intake Air Filter Assembly, Plastic Housing (1/4" x 1/4" Tube Fitting, NPT)'),
        new ItemCollection('TER1184120', 'ARB Compressor Under Seat Mounting Kit - 1184120', 'http://n3.datasn.io/data/api/v1/n3a2/auto_part_2/main/part_image//05/f9/5a/89/05f95a89fddc5411bb8d0a7dbc461fd8293111ba.jpg', 'The TeraFlex JK ARB compressor under seat mounting kit is a simple solution to mount the ARB on-board hi performance twin air compressor under the JK Unlimited front passenger seat and provides maximum protection from damage and from the elements...'),
    ];

    constructor(private http: HttpClient) {
    }

    get(code: string): Observable<ItemCollection> {
        return of(this.items.filter((col) => col.code === code)[0]);
    }

    getAll(): Observable<ItemCollection[]> {
        this.items.forEach(col => {
            col.items.length = 0;
        });
        return this.http.get<Item[]>(`${this.baseUrl}/api/v1/inventory`)
            .pipe(map((items) => {
                items.forEach(item => {
                    const index = this.items.findIndex((i) => i.code === item.code);
                    this.items[index].items.push(item);
                    this.items[index].manufacturer = item.manufacturer;
                    this.items[index].category = item.category;
                });
                return this.items;
            }));
    }

    add(item: Item): Observable<Item> {
        return this.http.post<Item>(`${this.baseUrl}/api/v1/inventory`, JSON.stringify(item));
    }

    transfer(item: Item): Observable<Item> {
        const req: UpdateRequest = {
            op: 'transfer',
            value: item.facility,
        };
        return this.http.put<Item>(`${this.baseUrl}/api/v1/inventory/${item.serial}`, JSON.stringify(req));
    }

    updateStatus(item: Item): Observable<Item> {
        const req: UpdateRequest = {
            op: 'status',
        };
        return this.http.put<Item>(`${this.baseUrl}/api/v1/inventory/${item.serial}`, JSON.stringify(req));
    }

    getHistory(item: Item): Observable<ItemHistory> {
        return this.http.get<ItemHistory>(`${this.baseUrl}/api/v1/inventory/${item.serial}/history`);
    }
}
