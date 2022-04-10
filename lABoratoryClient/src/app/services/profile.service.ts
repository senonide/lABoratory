import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ChartElement } from "../models/chart-element.model";
import { Customer } from "../models/customer.model";
import { Experiment } from "../models/experiment.model";
import { ExperimentService } from "./experiment.service";


@Injectable({providedIn: 'root'})
export class ProfileService {
    selectedExperiment: Experiment | null = null;
    theoreticalAssignments: any[] = [];
    formType: FormType = FormType.DEFAULT;

    actualAssignmentsCount: number = 0;
    actualAssignments: any[] = []

    constructor(private experimentService: ExperimentService) {}

    selectExperiment(experiment: Experiment): void {
        this.formType = FormType.EXPDET;
        this.selectedExperiment = experiment;
        var aux: any[] = [];
        for (let assignment of experiment.assignments){
            aux.push(ChartElement.getChartElementFromAssignment(assignment));
        }
        this.theoreticalAssignments = aux;
        this.generateActualAssignmentsData(experiment);
    }

    generateActualAssignmentsData(experiment: Experiment) {
        var assignments: Observable<Customer[]> | null = this.experimentService.getActualAssignments(experiment.id)
        if (assignments==null) {
            return;
        } else {
            assignments.subscribe({
                next: (assignments) => {
                    this.actualAssignmentsCount= assignments.length;
                    this.actualAssignments = ChartElement.getChartElementFromCustomer(assignments);
                },
                error: () => {}
            });
        }
    }
}

export enum FormType {
    DEFAULT,
    NEWEXP,
    UPDATE,
    EXPDET
}