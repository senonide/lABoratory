import { Component } from "@angular/core";
import { MatDialog } from "@angular/material/dialog";
import { Router } from "@angular/router";
import { Color, ScaleType } from "@swimlane/ngx-charts";
import { Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";
import { DeleteDialog } from "../delete-dialog/delete-dialog.component";
import { KeyDialog } from "../key-dialog/key-dialog.component";

@Component({
    selector: 'experiment-details',
    templateUrl: './experiment-details.component.html',
    styleUrls: ['./experiment-details.component.css']
})
export class ExperimentDetails {

    colorScheme: Color = { 
        domain: ['#54C6EB', '#FF3C38', '#FFBC42', '#69DC9E', '#6F58C9'], 
        group: ScaleType.Ordinal, 
        selectable: false, 
        name: 'Customer Usage', 
    };

    constructor(public profileService: ProfileService, private experimentService: ExperimentService, private router: Router, public dialog: MatDialog) {}

    openKeyDialog(): void {
        const dialogRef = this.dialog.open(KeyDialog, {
            data: {
                title: "Experiment key: ", 
                content: this.profileService.selectedExperiment?.experimentKey
            },
        });
    }

    openDeleteDialog(): void {
        const dialogRef = this.dialog.open(DeleteDialog, {
            data: {
                title: "Are you sure you want to delete '" + this.profileService.selectedExperiment?.name + "' ?", 
                content: " - All previous assignments will be removed as well."
            },
        });
    }

    updateExperiment(): void {
        this.profileService.formType = FormType.UPDATE;
    }

}