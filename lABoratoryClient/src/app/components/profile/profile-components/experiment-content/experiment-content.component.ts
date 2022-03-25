import { Component } from "@angular/core";
import { ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'experiment-content',
    templateUrl: './experiment-content.component.html',
    styleUrls: ['./experiment-content.component.css']
})
export class ExperimentContent {
    constructor(public profileService: ProfileService) {}
}