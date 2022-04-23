import { Component, OnInit } from "@angular/core";
import { FormGroup, FormControl, Validators, FormBuilder, FormArray } from "@angular/forms";
import { Router } from "@angular/router";
import { Assignment, Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'create-experiment',
    templateUrl: './create-experiment.component.html',
    styleUrls: ['./create-experiment.component.css']
})
export class CreateExperiment implements OnInit {

    newExperimentForm!: FormGroup;
    newExperiment!: Experiment;
    creating: boolean = false

    constructor(private experimentService: ExperimentService, public profileService: ProfileService, private router: Router, private formBuilder: FormBuilder) {}

    ngOnInit(): void {
        this.createForm();
    }

    createExperiment() {
        var newExperimentAssignments: any[] = [];
        newExperimentAssignments.push({
            assignmentValue: Number(this.newExperimentForm.value.controlAssignmentValue),
            assignmentDescription: "Control group"
        });
        var index: number = 1;
        for(let assignment of this.newExperimentForm.value.assignments) {
            if (assignment!=""){
                newExperimentAssignments.push({
                    assignmentValue: Number(assignment),
                    assignmentDescription: this.newExperimentForm.value.assignmentsDescriptions[index - 1]
                });
                index++;
            }
        }
        var newExperiment: Experiment = {
            id: "",
            name: this.newExperimentForm.value.name,
            description: this.newExperimentForm.value.description,
            experimentKey: "",
            assignments: newExperimentAssignments  
        };
        if(!this.validateExperiment(newExperiment)) return;
        this.creating = true;
        this.experimentService.createExperiment(newExperiment)?.subscribe({
            next: () => {
                this.newExperimentForm.reset();
                this.experimentService.getExperiments()?.subscribe({
                    next: (experiments) => {
                        this.experimentService.experiments = experiments;
                        this.createForm();
                        this.profileService.selectExperiment(this.experimentService.experiments[this.experimentService.experiments.length-1])
                        this.creating = false;
                    },
                    error: () => {
                        this.creating = false;
                        this.router.navigate(['/auth/login']);
                    }
                });
            },
            error: () => {
                this.creating = false;
            }
        });
    }

    private validateExperiment(experiment: Experiment): boolean {
        var acc: number = 0;
        for(let assignment of experiment.assignments) {
            acc += assignment.assignmentValue;
        }
        if(Math.round(acc) == 100) {
            return true;
        } else {
            return false;
        }
    }

    get assignments() {
        return this.newExperimentForm.get("assignments") as FormArray;
    }

    get assignmentsDescriptions() {
        return this.newExperimentForm.get("assignmentsDescriptions") as FormArray;
    }

    addAssignments() {
        this.assignments.push(
            this.formBuilder.control('', [Validators.required])
        );
        this.assignmentsDescriptions.push(
            this.formBuilder.control('')
        );
    }

    removeAssignment() {
        this.assignments.removeAt(this.assignments.length - 1)
        this.assignmentsDescriptions.removeAt(this.assignments.length - 1)
    }

    private createForm(): void {
        this.newExperimentForm = this.formBuilder.group({
            name: new FormControl('', [Validators.required]),
            description: new FormControl(''),
            controlAssignmentValue: new FormControl('', [Validators.required]),
            assignments: this.formBuilder.array([]),
            assignmentsDescriptions: this.formBuilder.array([])
        });
    }

}