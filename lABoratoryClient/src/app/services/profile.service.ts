import { Injectable } from "@angular/core";
import { ChartElement } from "../models/chart-element.model";
import { Experiment } from "../models/experiment.model";


@Injectable({providedIn: 'root'})
export class ProfileService {
    selectedExperiment: Experiment | null = null;
    theoreticalAssignments: any[] = [];
    formType: FormType = FormType.DEFAULT;

    actualAssignments: any[] = []

    selectExperiment(experiment: Experiment): void {
        this.formType = FormType.EXPDET;
        this.selectedExperiment = experiment;
        var aux: any[] = [];
        for (let assignment of experiment.assignments){
            aux.push(ChartElement.getChartElementFromAssignment(assignment));
        }
        this.theoreticalAssignments = aux;
    }

}

export enum FormType {
    DEFAULT,
    NEWEXP,
    UPDATE,
    EXPDET
}