import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'web-example',
  templateUrl: 'web-example.component.html',
  styleUrls: ['web-example.component.css']
})
export class WebExample {

  public price: String = "0"

  usernameForm = new FormGroup({
    username: new FormControl('', [Validators.required]),
  });

  
}
