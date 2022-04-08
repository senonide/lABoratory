import { Component, Inject } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";

import { Config } from "../../../../config/config";

@Component({
    selector: 'key-dialog',
    templateUrl: './key-dialog.component.html',
    styleUrls: ['./key-dialog.component.css']
})
export class KeyDialog {

    copied: boolean = false

    url: string = "";

    constructor(
        public dialogRef: MatDialogRef<KeyDialog>,
        @Inject(MAT_DIALOG_DATA) public data: DialogData) {
            this.url = Config.apiUrl + "/assignment/" + data.content
        }

        
}

export interface DialogData {
    title: string;
    content: string;
}