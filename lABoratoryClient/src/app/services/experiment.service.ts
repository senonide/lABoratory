import { Injectable } from "@angular/core";
import { HttpClient, HttpClientModule } from "@angular/common/http";
import { Subject } from "rxjs";
import { map } from "rxjs/operators";

import { Assignment, Experiment } from "../models/experiment.model";

@Injectable({providedIn: 'root'})
export class ExperimentService {

    private experiments: Experiment[] = [];

    private experimentsUpdated = new Subject<Experiment[]>();

    constructor(private http: HttpClient) {}

    getExperiments(){
        this.http
        .get<{message: string, posts: any}>(
            'http://localhost:8080/experiments'
        )
        .pipe(map((postData) => {
            return postData.posts.map((experiment: any) => {
                return {
                    id: experiment._id,
                    name: experiment.name,
                    assignments: experiment.assignments
                }
            });
        }))
        .subscribe(transformedExperiments => {
            this.experiments = transformedExperiments;
            //                     Unrefered copy
            this.experimentsUpdated.next([...this.experiments]);
        });
    }


    getExperimentUpdateListener(){
        return this.experimentsUpdated.asObservable();
    }



    addExperiment(name: string, assignments: Assignment[]){
        const experiment: Experiment = {id: '', name: name, assignments: assignments};
        this.http.post<{message: string, experimentId: string}>('http://localhost:8080/experiments', experiment)
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
        this.http.delete("http://localhost:8080/experiments" + experimentId)
        .subscribe(() => {
            const updatedExperiments = this.experiments.filter(experiment => experiment.id !== experimentId);
            this.experiments = updatedExperiments;
            // Reload the page
            this.experimentsUpdated.next([...this.experiments]);
        });
    }

}