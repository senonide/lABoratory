import { Component, OnInit } from "@angular/core";
import { FormArray, FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";
import { Router } from "@angular/router";
import { Color, ScaleType } from "@swimlane/ngx-charts";
import { Assignment, Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'experiment-content',
    templateUrl: './experiment-content.component.html',
    styleUrls: ['./experiment-content.component.css']
})
export class ExperimentContent implements OnInit {

    newExperimentForm!: FormGroup;
    newExperiment!: Experiment;

    colorScheme: Color = { 
        domain: ['#FF3C38', '#54C6EB',  '#FFBC42', '#69DC9E', '#6F58C9'], 
        group: ScaleType.Ordinal, 
        selectable: false, 
        name: 'Customer Usage', 
    };

    constructor(private experimentService: ExperimentService, public profileService: ProfileService, private router: Router, private formBuilder: FormBuilder) {}

    ngOnInit(): void {
        this.newExperimentForm = this.formBuilder.group({
            name: new FormControl('', [Validators.required]),
            controlAssignmentValue: new FormControl('', [Validators.required]),
            assignments: this.formBuilder.array([])
        });
    }

    deleteExperiment(experiment: Experiment) {
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
            error: () => {}
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
            assignments: newExperimentAssignments  
        };
        this.experimentService.createExperiment(newExperiment)?.subscribe({
            next: () => {
                this.newExperimentForm.reset();
                this.profileService.formType = FormType.DEFAULT
                console.log(newExperiment);
                this.experimentService.getExperiments()?.subscribe({
                    next: (experiments) => {
                        this.experimentService.experiments = experiments;
                        this.newExperimentForm = this.formBuilder.group({
                            name: new FormControl('', [Validators.required]),
                            controlAssignmentValue: new FormControl('', [Validators.required]),
                            assignments: this.formBuilder.array([])
                        });
                    },
                    error: () => {
                        this.router.navigate(['/auth/login']);
                    }
                });
            },
            error: () => {}
        });
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