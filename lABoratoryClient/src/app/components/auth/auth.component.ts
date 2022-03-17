import {Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

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

    public hasError: boolean = false;
    public errorMessage: string = '';

    username = new FormControl('', [Validators.required, Validators.email]);
    password = new FormControl('', [Validators.required, Validators.minLength(1)]);
    repeatedPassword =  new FormControl('', [Validators.required, Validators.minLength(1)]);

    public constructor(private route:ActivatedRoute, private authService: AuthService, private router: Router){}

    authenticate(form: NgForm) {
        if (form.invalid) return;
        this.hasError = false;
        this.authService.login(form.value.username, form.value.password).subscribe({
            next: () => {
                form.resetForm();
                this.router.navigate(['/profile']);
            },
            error: () => {
                this.hasError = true;
                this.errorMessage = 'Invalid user or password supplied!';
                console.log(this.errorMessage);
            }
        });
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