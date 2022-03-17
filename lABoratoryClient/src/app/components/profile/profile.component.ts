import {Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';

import { ExperimentService } from 'src/app/services/experiment.service';
import { Experiment } from 'src/app/models/experiment.model';

@Component({
    selector: 'profile-component',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {

    experiments: Experiment[] = [];
    private experimentsSub: Subscription = new Subscription;

    constructor(private experimentService: ExperimentService) {}
    
    ngOnInit(): void {
        this.experimentService.getExperiments();
        this.experimentsSub = this.experimentService.getExperimentUpdateListener().subscribe(
            (experiments: Experiment[])=>{this.experiments = experiments;}
        );
    }

    ngOnDestroy(): void {
        this.experimentsSub.unsubscribe();
    }

}