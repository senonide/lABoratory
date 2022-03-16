export interface Experiment {
    id: string;
    name: string;
    assignments: Assignment[];
}

export interface Assignment {
    assignmentName: string;
    assignmentValue: number;

}