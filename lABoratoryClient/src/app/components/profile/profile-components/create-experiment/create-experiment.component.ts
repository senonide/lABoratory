import { Component, OnInit } from "@angular/core";
import { FormGroup, FormControl, Validators, FormBuilder, FormArray } from "@angular/forms";
import { Router } from "@angular/router";
import { ChartElement } from "src/app/models/chart-element.model";
import { Assignment, Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";

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
        this.newExperimentForm = this.formBuilder.group({
            name: new FormControl('', [Validators.required]),
            description: new FormControl(''),
            controlAssignmentValue: new FormControl('', [Validators.required]),
            assignments: this.formBuilder.array([])
        });
    }

    createExperiment() {
        var newExperimentAssignments: Assignment[] = [];
        newExperimentAssignments.push({
            assignmentName: "c",
            assignmentValue: Number(this.newExperimentForm.value.controlAssignmentValue)
        });
        var index: number = 1;
        for(let assignment of this.newExperimentForm.value.assignments) {
            if (assignment!=""){
                newExperimentAssignments.push({
                    assignmentName: "a" + index,
                    assignmentValue: Number(assignment)
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
                        this.newExperimentForm = this.formBuilder.group({
                            name: new FormControl('', [Validators.required]),
                            description: new FormControl(''),
                            controlAssignmentValue: new FormControl('', [Validators.required]),
                            assignments: this.formBuilder.array([])
                        });
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

    addAssignments() {
        this.assignments.push(
            this.formBuilder.control('')
        );
    }

}