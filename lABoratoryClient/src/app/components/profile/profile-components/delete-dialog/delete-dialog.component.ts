import { Component, Inject } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { Router } from "@angular/router";
import { Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'delete-dialog',
    templateUrl: './delete-dialog.component.html',
    styleUrls: ['./delete-dialog.component.css']
})
export class DeleteDialog {

    constructor(
        public profileService: ProfileService, 
        private experimentService: ExperimentService,
        private router: Router,
        public dialogRef: MatDialogRef<DeleteDialog>,
        @Inject(MAT_DIALOG_DATA) public data: DialogData,) {}

    onNoClick(): void {
        this.dialogRef.close();
    }

    deleteExperiment(experiment: Experiment | null) {
        if (experiment == null){
            return
        }
        this.experimentService.deleteExperiment(experiment)?.subscribe({
            next: () => {
                this.experimentService.getExperiments()?.subscribe({
                    next: (experiments) => {
                        this.experimentService.experiments = experiments;
                        this.profileService.formType = FormType.DEFAULT
                    },
                    error: () => {
                        this.router.navigate(['/auth/login']);
                    }
                });
            },
            error: () => {
            }
        });
        this.dialogRef.close();
    }
}

export interface DialogData {
    title: string;
    content: string;
}