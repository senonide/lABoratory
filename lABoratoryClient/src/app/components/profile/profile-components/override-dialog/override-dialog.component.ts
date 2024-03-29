import { Component, Inject } from "@angular/core";
import { FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { Router } from "@angular/router";
import { Customer } from "src/app/models/customer.model";
import { Experiment, Assignment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'override-dialog',
    templateUrl: './override-dialog.component.html',
    styleUrls: ['./override-dialog.component.css']
})
export class OverrideDialog {

    public overrides: Customer[] = [];

    overrideForm: FormGroup =  this.formBuilder.group({
        key: new FormControl('', [Validators.required]),
    });;

    constructor(
        public dialogRef: MatDialogRef<OverrideDialog>,
        private experimentService: ExperimentService,
        private router: Router,
        private profileService: ProfileService,
        private formBuilder: FormBuilder,
        @Inject(MAT_DIALOG_DATA) public data: DialogData) {
            var response = this.experimentService.getOverrideAssignments(data.experiment.id)
            if(response==null){
                return
            } else {
                response.subscribe({
                    next: (overrides) => {
                        this.overrides = overrides;
                    },
                    error: () => {
                        this.router.navigate(['/auth/login']);
                    }
                });
            }
    }

    overrideCustomer(experiment: Experiment, assignment: Assignment): void {
        var response  = this.experimentService.overrideCustomer(experiment.experimentKey, this.overrideForm.value.key, assignment.assignmentName);
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

    deleteOverride(experiment: string, key: string): void {
        var response  = this.experimentService.deleteOverride(experiment, key);
        if(response==null){
            return
        } else {
            response.subscribe({
                next: () => {
                    this.dialogRef.close();
                },
                error: () => {
                    this.dialogRef.close();
                    //this.router.navigate(['/auth/login']);
                }
            });
        }
    }

}

export interface DialogData {
    title: string;
    experiment: Experiment;
}