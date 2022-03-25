import { Component } from "@angular/core";
import { Router } from "@angular/router";
import { Color, ScaleType } from "@swimlane/ngx-charts";
import { Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'experiment-details',
    templateUrl: './experiment-details.component.html',
    styleUrls: ['./experiment-details.component.css']
})
export class ExperimentDetails {

    colorScheme: Color = { 
        domain: ['#FF3C38', '#54C6EB',  '#FFBC42', '#69DC9E', '#6F58C9'], 
        group: ScaleType.Ordinal, 
        selectable: false, 
        name: 'Customer Usage', 
    };

    constructor(public profileService: ProfileService, private experimentService: ExperimentService, private router: Router) {}

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
}