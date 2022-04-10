import { Customer } from "./customer.model";
import { Assignment, Experiment } from "./experiment.model";

export class ChartElement {
    public static getChartElementFromAssignment( assignment: Assignment ) {
        return {
            name: assignment.assignmentName.toUpperCase(),
            value: assignment.assignmentValue
        };
    }
    public static getChartElementFromCustomer(customers: Customer[]): any[] {
        let data = new Map<string, number>();
        for (let customer of customers) {
            let count = data.get(customer.assignment)
            let aux = 1;
            if (count!=null) {
                aux = count + 1;
            }
            data.set(customer.assignment, aux);
        }
        var result: any[] = []
        for (let [key, value] of data) {
            result.push({
                name: key,
                value: value
            })
        }
        return result;
    }
}