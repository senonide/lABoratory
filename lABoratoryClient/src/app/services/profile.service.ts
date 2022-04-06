import { Injectable } from "@angular/core";
import { Experiment } from "../models/experiment.model";


@Injectable({providedIn: 'root'})
export class ProfileService {
    selectedExperiment: Experiment | null = null;
    currentAssignments: any[] = [];
    formType: FormType = FormType.DEFAULT;
}

export enum FormType {
    DEFAULT,
    NEWEXP,
    EXPDET
}