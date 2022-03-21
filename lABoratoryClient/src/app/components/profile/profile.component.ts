import {Component, OnDestroy, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

import { ExperimentService } from 'src/app/services/experiment.service';
import { Experiment } from 'src/app/models/experiment.model';
import { Router } from '@angular/router';

@Component({
    selector: 'profile-component',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {

    experiments: Experiment[] = [];

    constructor(private experimentService: ExperimentService, private router: Router) {}

    logout(): void {
        localStorage.removeItem('jwt');
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