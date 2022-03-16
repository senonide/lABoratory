import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Router } from "@angular/router";

@Injectable({providedIn: 'root'})
export class AuthService {

    constructor(private http: HttpClient, private router: Router) {}

    authUser(usernameInForm: string, passwordInForm: string, type: string) {
        const credentials = {username: usernameInForm, password: passwordInForm };
        var token: string = '';
        var url = '';
        if(type==='login'){
            url = 'http://localhost:8080/auth';
        } else if(type==='signup'){
            url = 'http://localhost:8080/signup';
        }
        this.http.post<{response: string, token: string}>(url, credentials)
        .subscribe((responseData) => {
                console.log(responseData.response + " " + responseData.token);
                token = responseData.token;
                if(token!=''){
                    localStorage.setItem('jwt', token);
                    this.router.navigate(['/profile'])
                }
            }
        );
    }
}