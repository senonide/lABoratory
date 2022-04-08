import { Component, OnInit } from "@angular/core";
import { ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'update-experiment',
    templateUrl: './update-experiment.component.html',
    styleUrls: ['./update-experiment.component.css']
})
export class UpdateExperiment implements OnInit {

    constructor(public profileService: ProfileService) {}

    ngOnInit(): void {
    }
    
}