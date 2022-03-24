import {Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable } from 'rxjs';

import { FormArray, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

import { ExperimentService } from 'src/app/services/experiment.service';
import { Assignment, Experiment } from 'src/app/models/experiment.model';
import { Router } from '@angular/router';

import { Color, ScaleType } from '@swimlane/ngx-charts';
import { ChartElement } from 'src/app/models/chart-element.model';

export enum FormType {
    DEFAULT,
    NEWEXP,
    EXPDET
}

@Component({
    selector: 'profile-component',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {

    public experiments: Experiment[] = [];

    newExperimentForm!: FormGroup;
    newExperiment!: Experiment;

    selectedExperiment!: Experiment;
    currentAssignments: any[] = [];
    formType: FormType = FormType.DEFAULT;

    colorScheme: Color = { 
        domain: ['#FF3C38', '#54C6EB',  '#FFBC42', '#69DC9E', '#6F58C9'], 
        group: ScaleType.Ordinal, 
        selectable: false, 
        name: 'Customer Usage', 
    };

    constructor(private experimentService: ExperimentService, private router: Router, private formBuilder: FormBuilder) {
    }

    logout(): void {
        localStorage.removeItem('jwt');
    }

    newExperimentOption() {
        this.formType = FormType.NEWEXP
    }

    selectExperiment(experiment: Experiment): void {
        this.formType = FormType.EXPDET;
        this.selectedExperiment = experiment;
        var aux: any[] = [];
        for (let assignment of experiment.assignments){
            aux.push(ChartElement.getChartElementFromAssignment(assignment));
        }
        this.currentAssignments = aux;
       
    }

    get assignments() {
        return this.newExperimentForm.get("assignments") as FormArray;
    }

    addAssignments() {
        this.assignments.push(
            this.formBuilder.control('')
        );
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
                this.formType = FormType.DEFAULT
                console.log(newExperiment);
                this.experimentService.getExperiments()?.subscribe({
                    next: (experiments) => {
                        this.experiments = experiments;
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
            error: () => {
                
            }
        });
        
    }

    deleteExperiment(experiment: Experiment) {
        this.experimentService.deleteExperiment(experiment)?.subscribe({
            next: () => {
                this.experimentService.getExperiments()?.subscribe({
                    next: (experiments) => {
                        this.experiments = experiments;
                        this.formType = FormType.DEFAULT
                    },
                    error: () => {
                        this.router.navigate(['/auth/login']);
                    }
                });
            },
            error: () => {
                
            }
        });
    }
    
    ngOnInit(): void {
        var exp: Observable<Experiment[]> | null = this.experimentService.getExperiments();
        if(exp==null) {
            this.router.navigate(['/auth/login']);
        } else {
            exp.subscribe({
                next: (experiments) => {
                    this.experiments = experiments;
                },
                error: () => {
                    this.router.navigate(['/auth/login']);
                }
            });
        }

        this.newExperimentForm = this.formBuilder.group({
            name: new FormControl('', [Validators.required]),
            controlAssignmentValue: new FormControl('', [Validators.required]),
            assignments: this.formBuilder.array([])
        });
    }

    ngOnDestroy(): void {
       
    }

}