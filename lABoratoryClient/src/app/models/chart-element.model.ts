import { Assignment } from "./experiment.model";

export class ChartElement {
    public static getChartElementFromAssignment( assignment: Assignment ) {
        return {
            name: assignment.assignmentName.toUpperCase(),
            value: assignment.assignmentValue
        };
    }
}