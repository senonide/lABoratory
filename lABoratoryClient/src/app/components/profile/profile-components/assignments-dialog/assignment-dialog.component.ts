import { Component, Inject } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { Router } from "@angular/router";
import { Experiment, Assignment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'assignment-dialog',
    templateUrl: './assignment-dialog.component.html',
    styleUrls: ['./assignment-dialog.component.css']
})
export class AssignmentDialog {

    copied: boolean = false

    constructor(
        public dialogRef: MatDialogRef<AssignmentDialog>,
        private experimentService: ExperimentService,
        private router: Router,
        private profileService: ProfileService,
        @Inject(MAT_DIALOG_DATA) public data: DialogData) {        
    }

    activateOneAssignment(experiment: Experiment, assignment: Assignment): void {
        for(let a of experiment.assignments) {
            if(a.assignmentName == assignment.assignmentName){
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
                    this.dialogRef.close();
                },
                error: () => {
                }
            });
        }
    }

}

export interface DialogData {
    title: string;
    experiment: Experiment;
}