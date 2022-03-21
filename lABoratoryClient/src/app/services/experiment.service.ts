import { Injectable } from "@angular/core";
import { HttpClient, HttpClientModule, HttpHeaders } from "@angular/common/http";
import { Observable, observable, Subject } from "rxjs";
import { map } from "rxjs/operators";

import { Assignment, Experiment } from "../models/experiment.model";
import { Config } from "../config";

@Injectable({providedIn: 'root'})
export class ExperimentService {

    private experiments: Experiment[] = [];

    private jwt: string = '';

    constructor(private http: HttpClient) {}

    getExperiments() : Observable<Experiment[]> | null{
        var auxJwt: string | null = localStorage.getItem('jwt');
        if (auxJwt!= null){
            this.jwt = auxJwt;
        } else {
            return null;
        }
        const httpOptions = {
            headers: new HttpHeaders({
                'Authorization':  this.jwt,
            })
        };
        return this.http.get<Experiment[]>(Config.apiUrl + '/experiments', httpOptions);
    }

/*
    addExperiment(name: string, assignments: Assignment[]){
        const experiment: Experiment = {id: '', name: name, assignments: assignments};
        this.http.post<{message: string, experimentId: string}>(Config.apiUrl + '/experiments', experiment)
        .subscribe(responseData => {
                console.log(responseData.message);
                const id = responseData.experimentId;
                experiment.id = id;
                this.experiments.push(experiment);
                this.experimentsUpdated.next([...this.experiments]);
            }
        );
        
    }

    deleteExperiment(experimentId: string) {
        this.http.delete(Config.apiUrl + '/experiments' + experimentId)
        .subscribe(() => {
            const updatedExperiments = this.experiments.filter(experiment => experiment.id !== experimentId);
            this.experiments = updatedExperiments;
            // Reload the page
            this.experimentsUpdated.next([...this.experiments]);
        });
    }
*/
}