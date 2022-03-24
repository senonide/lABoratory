import {Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import {FormControl, Validators, FormGroup} from '@angular/forms';

import { AuthService } from 'src/app/services/auth.service';

@Component({
    selector: 'auth-component',
    templateUrl: './auth.component.html',
    styleUrls: ['./auth.component.css'],
})
export class AuthComponent implements OnInit, OnDestroy {

    public type: any;
    public hasError: boolean = false;

    authForm = new FormGroup({
        username: new FormControl('', [Validators.required, Validators.email]),
        password: new FormControl('', [Validators.required, Validators.minLength(1)]),
    });

    signupForm = new FormGroup({
        username: new FormControl('', [Validators.required, Validators.email]),
        password: new FormControl('', [Validators.required, Validators.minLength(1)]),
        repeatedPassword: new FormControl('', [Validators.required, Validators.minLength(1)]),
    });

    public constructor(private route:ActivatedRoute, private authService: AuthService, private router: Router){}

    authenticate() {
        if (this.authForm.invalid) return;
        this.hasError = false;
        this.authService.login(this.authForm.value.username, this.authForm.value.password).subscribe({
            next: () => {
                this.authForm.reset();
                this.router.navigate(['/profile']);
            },
            error: () => {
                this.authForm.controls['password'].setErrors({
                    incorrect: true
                });
            }
        });
    }

    signup() {
        if (this.signupForm.invalid) return;
        if(this.signupForm.controls['password'].value != this.signupForm.controls['repeatedPassword'].value) {
            this.signupForm.controls['repeatedPassword'].setErrors({
                notMatch: true
            });
            return;
        } 
        this.hasError = false;
        this.authService.signup(this.signupForm.value.username, this.signupForm.value.password).subscribe({
            next: () => {
                this.signupForm.reset();
                this.router.navigate(['/profile']);
            },
            error: () => {
                this.authForm.controls['password'].setErrors({
                    incorrect: true
                });
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
