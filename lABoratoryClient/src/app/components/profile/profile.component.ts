import {Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { NgForm } from "@angular/forms";
import {FormControl, Validators} from '@angular/forms';

import { AuthService } from 'src/app/services/auth.service';
import { Injectable } from "@angular/core";
import { HttpClient, HttpClientModule, HttpHeaders, HttpResponse } from "@angular/common/http";
import { observable } from 'rxjs';


@Component({
    selector: 'profile-component',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {

    public jwt: string = '';

    constructor(private http: HttpClient) {}
    
    ngOnInit(): void {
        var auxJwt: string | null = localStorage.getItem('jwt');
        if(auxJwt!==null){
            this.jwt = auxJwt;
        }
        const httpOptions = {
            headers: new HttpHeaders({
                'Authorization':  this.jwt,
            })
        };
        this.http.get('http://localhost:8080/experiments', httpOptions)
        .subscribe((responseData) => {
                console.log(responseData);
            }
        );
    }

    ngOnDestroy(): void {
        
    }
    
}