export interface Experiment {
    id: string;
    name: string;
    description: string;
    experimentKey: string;
    assignments: Assignment[];
}

export interface Assignment {
    assignmentName: string;
    assignmentValue: number;
    assignmentDescription: string;
}