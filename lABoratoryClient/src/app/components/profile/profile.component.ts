import { Component } from '@angular/core';

import { FormType, ProfileService } from 'src/app/services/profile.service';

@Component({
    selector: 'profile-component',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

    constructor(public profileService: ProfileService) {}

    logout(): void {
        localStorage.removeItem('jwt');
        this.profileService.currentAssignments = [];
        this.profileService.formType = FormType.DEFAULT;
        this.profileService.selectedExperiment = null;
    }

    newExperimentOption() {
        this.profileService.isTryingToDelete = false;
        this.profileService.formType = FormType.NEWEXP;
    }
}