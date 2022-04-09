import { Component, OnInit } from "@angular/core";
import { FormArray, FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";
import { Router } from "@angular/router";
import { Assignment, Experiment } from "src/app/models/experiment.model";
import { ExperimentService } from "src/app/services/experiment.service";
import { FormType, ProfileService } from "src/app/services/profile.service";

@Component({
    selector: 'update-experiment',
    templateUrl: './update-experiment.component.html',
    styleUrls: ['./update-experiment.component.css']
})
export class UpdateExperiment implements OnInit {

    updateExperimentForm!: FormGroup;

    updating: boolean = false;

    constructor(private experimentService: ExperimentService, public profileService: ProfileService, private router: Router, private formBuilder: FormBuilder) {}

    ngOnInit(): void {
        this.updateExperimentForm = this.formBuilder.group({
            name: new FormControl(this.profileService.selectedExperiment?.name, [Validators.required]),
            description: new FormControl(this.profileService.selectedExperiment?.description),
            controlAssignmentValue: new FormControl(this.profileService.selectedExperiment?.assignments[0].assignmentValue, [Validators.required]),
            assignments: this.formBuilder.array(this.getAssignmentsFormControls()),
            assignmentsDescriptions: this.formBuilder.array(this.getAssignmentsDescriptionsFormControls())
        });
    }

    back(): void {
        this.profileService.formType = FormType.EXPDET;
    }

    updateExperiment(): void {
        var ExperimentAssignments: Assignment[] = [];
        ExperimentAssignments.push({
            assignmentName: "c",
            assignmentValue: Number(this.updateExperimentForm.value.controlAssignmentValue),
            assignmentDescription: "Control group"
        });
        var index: number = 1;
        for(let assignment of this.updateExperimentForm.value.assignments) {
            if (assignment!=""){
                ExperimentAssignments.push({
                    assignmentName: "a" + index,
                    assignmentValue: Number(assignment),
                    assignmentDescription: this.updateExperimentForm.value.assignmentsDescriptions[index - 1]
                });
                index++;
            }
        }
        if(this.profileService.selectedExperiment == null) {
            return;
        }
        var updatedExperiment: Experiment = {
            id: this.profileService.selectedExperiment?.id,
            name: this.updateExperimentForm.value.name,
            description: this.updateExperimentForm.value.description,
            experimentKey: "",
            assignments: ExperimentAssignments  
        };
        console.log(updatedExperiment);
        if(!this.validateExperiment(updatedExperiment)) return;
        this.updating = true;
        this.experimentService.updateExperiment(updatedExperiment)?.subscribe({
            next: () => {
                this.experimentService.getExperiments()?.subscribe({
                    next: (experiments) => {
                        this.experimentService.experiments = experiments;
                        this.updating = false;
                        this.profileService.selectExperiment(updatedExperiment);
                        this.back();
                    },
                    error: () => {
                        this.updating = false;
                        this.router.navigate(['/auth/login']);
                    }
                });
            },
            error: () => {
                this.updating = false;
            }
        });
    }

    private validateExperiment(experiment: Experiment): boolean {
        var acc: number = 0;
        for(let assignment of experiment.assignments) {
            acc += assignment.assignmentValue;
        }
        if(Math.round(acc) == 100) {
            return true;
        } else {
            return false;
        }
    }

    get assignments() {
        return this.updateExperimentForm.get("assignments") as FormArray;
    }

    getAssignmentsFormControls(): FormControl[]{
        if ( this.profileService.selectedExperiment?.assignments==null){
            return [];
        }
        var output: FormControl[] = []
        for (let assignment of this.profileService.selectedExperiment?.assignments.slice(1, this.profileService.selectedExperiment?.assignments.length)) {
            output.push(new FormControl(assignment.assignmentValue))
        }
        return output;
    }

    get assignmentsDescriptions() {
        return this.updateExperimentForm.get("assignmentsDescriptions") as FormArray;
    }

    getAssignmentsDescriptionsFormControls(){
        if ( this.profileService.selectedExperiment?.assignments==null){
            return [];
        }
        var output: FormControl[] = []
        for (let assignment of this.profileService.selectedExperiment?.assignments.slice(1, this.profileService.selectedExperiment?.assignments.length)) {
            output.push(new FormControl(assignment.assignmentDescription))
        }
        return output;
    }

    addAssignments() {
        this.assignments.push(
            this.formBuilder.control('', [Validators.required])
        );
        this.assignmentsDescriptions.push(
            this.formBuilder.control('')
        );
    }

    removeAssignment() {
        this.assignments.removeAt(this.assignments.length - 1)
        this.assignmentsDescriptions.removeAt(this.assignments.length - 1)
    }
    
}