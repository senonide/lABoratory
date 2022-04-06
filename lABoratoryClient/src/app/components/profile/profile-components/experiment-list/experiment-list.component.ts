import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { ChartElement } from "src/app/models/chart-element.model";
import { Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'experiment-list',
    templateUrl: './experiment-list.component.html',
    styleUrls: ['./experiment-list.component.css']
})
export class ExperimentList implements OnInit {

    constructor(public experimentService: ExperimentService, private profileService: ProfileService, private router: Router) {}

    ngOnInit(): void {
        var exp: Observable<Experiment[]> | null = this.experimentService.getExperiments();
        if(exp==null) {
            this.router.navigate(['/auth/login']);
        } else {
            exp.subscribe({
                next: (experiments) => {
                    this.experimentService.experiments = experiments;
                },
                error: () => {
                    this.router.navigate(['/auth/login']);
                }
            });
        }
    }

    selectExperiment(experiment: Experiment): void {
        this.profileService.formType = FormType.EXPDET;
        this.profileService.selectedExperiment = experiment;
        var aux: any[] = [];
        for (let assignment of experiment.assignments){
            aux.push(ChartElement.getChartElementFromAssignment(assignment));
        }
        this.profileService.currentAssignments = aux;
    }
}