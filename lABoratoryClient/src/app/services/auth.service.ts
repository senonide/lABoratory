import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Router } from "@angular/router";
import { map, Observable } from "rxjs";

export interface AuthResponse{
    response: string,
    token: string
}

@Injectable({providedIn: 'root'})
export class AuthService {

    constructor(private http: HttpClient) {}

    login(username: string, password: string): Observable<void> {
        const credentials = { username, password };
        return this.http.post<AuthResponse>('http://localhost:8080/auth', credentials)
            .pipe(map(response => {
                localStorage.setItem('jwt', response.token);
            }));
    }

    signup(username: string, password: string): Observable<void> {
        const credentials = { username, password };
        return this.http.post<AuthResponse>('http://localhost:8080/auth', credentials)
            .pipe(map(response => {
                localStorage.setItem('jwt', response.token);
            }));
    }
}