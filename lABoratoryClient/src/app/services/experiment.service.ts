import { Injectable } from "@angular/core";
import { HttpClient, HttpClientModule, HttpHeaders } from "@angular/common/http";
import { Observable, observable, Subject } from "rxjs";
import { map } from "rxjs/operators";

import { Assignment, Experiment } from "../models/experiment.model";
import { Config } from "../config/config";
import { Customer } from "../models/customer.model";

@Injectable({providedIn: 'root'})
export class ExperimentService {

    public experiments: Experiment[] = [];

    private jwt: string = '';

    constructor(private http: HttpClient) {}

    getExperiments(){
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

    deleteExperiment(experiment: Experiment) {
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
        return this.http.delete(Config.apiUrl + '/experiments/' + experiment.id, httpOptions);
    }

    createExperiment(experiment: Experiment) {
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
        return this.http.post(Config.apiUrl + '/experiments', experiment, httpOptions);
    }

    updateExperiment(experiment: Experiment) {
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
        return this.http.put(Config.apiUrl + '/experiments/' + experiment.id, experiment, httpOptions);
    }

    getActualAssignments(experimentId: string) {
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
        return this.http.get<Customer[]>(Config.apiUrl + '/assignments/' + experimentId, httpOptions);
    }

}