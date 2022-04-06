import { Component, Inject } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";

@Component({
    selector: 'key-dialog',
    templateUrl: './key-dialog.component.html',
    styleUrls: ['./key-dialog.component.css']
})
export class KeyDialog {
    constructor(
        public dialogRef: MatDialogRef<KeyDialog>,
        @Inject(MAT_DIALOG_DATA) public data: DialogData,) {}
}

export interface DialogData {
    title: string;
    content: string;
}