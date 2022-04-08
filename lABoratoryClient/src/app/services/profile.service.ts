import { Injectable } from "@angular/core";
import { ChartElement } from "../models/chart-element.model";
import { Experiment } from "../models/experiment.model";


@Injectable({providedIn: 'root'})
export class ProfileService {
    selectedExperiment: Experiment | null = null;
    currentAssignments: any[] = [];
    formType: FormType = FormType.DEFAULT;


    selectExperiment(experiment: Experiment): void {
        this.formType = FormType.EXPDET;
        this.selectedExperiment = experiment;
        var aux: any[] = [];
        for (let assignment of experiment.assignments){
            aux.push(ChartElement.getChartElementFromAssignment(assignment));
        }
        this.currentAssignments = aux;
    }

}

export enum FormType {
    DEFAULT,
    NEWEXP,
    UPDATE,
    EXPDET
}