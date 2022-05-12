import { Component } from "@angular/core";
import { MatDialog } from "@angular/material/dialog";
import { Router } from "@angular/router";
import { Color, ScaleType } from "@swimlane/ngx-charts";
import { Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";
import { AssignmentDialog } from "../assignments-dialog/assignment-dialog.component";
import { DeleteDialog } from "../delete-dialog/delete-dialog.component";
import { KeyDialog } from "../key-dialog/key-dialog.component";
import { OverrideDialog } from "../override-dialog/override-dialog.component";

@Component({
    selector: 'experiment-details',
    templateUrl: './experiment-details.component.html',
    styleUrls: ['./experiment-details.component.css']
})
export class ExperimentDetails {

    colorScheme1: Color = { 
        domain: ['#1666ba', '#368ce7', '#7ab3ef', '#bedaf7', '#deecfb'], 
        group: ScaleType.Ordinal, 
        selectable: false, 
        name: 'Customer Usage', 
    };

    colorScheme2: Color = { 
        domain: ['#991101', '#c33211', '#d75f5a', '#ff8a82', '#ffcfc2'], 
        group: ScaleType.Ordinal, 
        selectable: false, 
        name: 'Customer Usage', 
    };

    constructor(public profileService: ProfileService, private experimentService: ExperimentService, private router: Router, public dialog: MatDialog) {}

    openKeyDialog(): void {
        const dialogRef = this.dialog.open(KeyDialog, {
            //width: "70%",
            data: {
                title: "Experiment key: ", 
                experimentService: this.experimentService,
                router: this.router,
                profileService: this.profileService,
                content: this.profileService.selectedExperiment?.experimentKey
            },
        });
    }

    openAssignmentsDialog(): void {
        const dialogRef = this.dialog.open(AssignmentDialog, {
            //width: "70%",
            data: {
                title: "Choose one assignment to activate: ", 
                experiment: this.profileService.selectedExperiment!,
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

    disableExperiment(experiment: Experiment): void {
        for(let a of experiment.assignments) {
            if(a.assignmentName == "c"){
                a.assignmentValue = 100.0;
            } else {
                a.assignmentValue = 0.0;
            }
        }
        var response = this.experimentService.updateExperiment(experiment)
        if(response==null){
            return
        } else {
            response.subscribe({
                next: () => {
                    this.experimentService.getExperiments()?.subscribe({
                        next: (experiments) => {
                            this.experimentService.experiments = experiments;
                        },
                        error: () => {
                            this.router.navigate(['/auth/login']);
                        }
                    });
                    this.profileService.selectExperiment(experiment);
                },
                error: () => {
                }
            });
        }
    }

    overrideCustomer(): void {
        const dialogRef = this.dialog.open(OverrideDialog, {
            data: {
                title: "Override a customer's assignment: ", 
                experiment: this.profileService.selectedExperiment!,
            },
        });
    }

    updateExperiment(): void {
        this.profileService.formType = FormType.UPDATE;
    }

}