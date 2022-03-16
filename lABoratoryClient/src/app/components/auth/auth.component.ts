import {Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { NgForm } from "@angular/forms";
import {FormControl, Validators} from '@angular/forms';

import { AuthService } from 'src/app/services/auth.service';


@Component({
    selector: 'auth-component',
    templateUrl: './auth.component.html',
    styleUrls: ['./auth.component.css']
})
export class AuthComponent implements OnInit, OnDestroy {

    public type: any;

    username = new FormControl('', [Validators.required, Validators.email]);
    password = new FormControl('', [Validators.required, Validators.minLength(1)]);
    repeatedPassword =  new FormControl('', [Validators.required, Validators.minLength(1)]);
    public constructor(private route:ActivatedRoute, private authService: AuthService){}

    authenticate(form: NgForm){
        if(form.invalid) return;
        this.authService.authUser(form.value.username, form.value.password, this.type);
        form.resetForm();
    }

    public ngOnInit(): void {
        this.route.paramMap.subscribe((paramMap: any) => {
            const {params} = paramMap;
            this.type = params.type;
        })
        
    }

    public ngOnDestroy(): void {
        
    }

}