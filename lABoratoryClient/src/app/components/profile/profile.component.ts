import {Component, OnDestroy, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

import { FormArray } from '@angular/forms';

import { ExperimentService } from 'src/app/services/experiment.service';
import { Experiment } from 'src/app/models/experiment.model';
import { Router } from '@angular/router';
import { NgForm } from '@angular/forms';


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

    experiments: Experiment[] = [];
    selectedExperiment!: Experiment;
    formType: FormType = FormType.DEFAULT;

    constructor(private experimentService: ExperimentService, private router: Router) {}

    logout(): void {
        localStorage.removeItem('jwt');
    }

    newExperimentOption() {
        this.formType = FormType.NEWEXP
    }

    selectExperiment(experiment: Experiment): void {
        this.formType = FormType.EXPDET;
        this.selectedExperiment = experiment;
    }

    newExperiment(form: NgForm) {

    }

    addAssignment() {
        
    }

    editExpreiment() {

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
    }

    ngOnDestroy(): void {
       
    }

}