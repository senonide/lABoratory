import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Config } from 'src/app/config/config';
import { Customer } from 'src/app/models/customer.model';

@Component({
  selector: 'web-example',
  templateUrl: 'web-example.component.html',
  styleUrls: ['web-example.component.css']
})
export class WebExample {

  public price: String = "0";

  public experimentId: String = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MjgzNDg2OGFhNmE1M2U2YzJlZTJmOTMifQ.xpBvvZd1S8Dz08SKD-WxGKlIIMft6fQv5k2B-D2okQQ";

  public assignment: String = "";

  public logged: boolean = false;

  constructor(private http: HttpClient) {}

  usernameForm = new FormGroup({
    username: new FormControl('', [Validators.required]),
  });

  loadNewUser(): void {
    var experimentKey: String = this.usernameForm.value.username;
    var response = this.getAssignment(this.experimentId, experimentKey);
      if(response==null){
          return
      } else {
        response.subscribe({
            next: (value) => {
                this.assignment = value.assignment;
                this.getPriceFromAssignment(this.assignment);
                this.logged = true;
            },
            error: () => {
              console.log("ERROR")
            }
        });
      }
  }

  getAssignment(experimentId: String, experimentKey: String) {
    return this.http.get<Customer>(Config.apiUrl + '/assignment' + '/' + experimentId + '/' + experimentKey);
  }

  getPriceFromAssignment(assignment: String): void {
    switch (assignment) {
      case "c":
        this.price = "500";
        break;
      case "a1":
        this.price = "499";
        break;
      case "a2":
        this.price = "450";
        break;
      case "a3":
        this.price = "489";
        break;
      case "a4":
        this.price = "510";
        break;
    }
  }

  

  
}
