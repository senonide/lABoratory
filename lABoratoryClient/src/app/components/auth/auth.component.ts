import {Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
    selector: 'auth-component',
    templateUrl: './auth.component.html',
    styleUrls: ['./auth.component.css']
})
export class AuthComponent implements OnInit, OnDestroy {

    public type: any;

    public constructor(private route:ActivatedRoute){}

    public ngOnInit(): void {
        this.route.paramMap.subscribe((paramMap: any) => {
            const {params} = paramMap;
            this.type = params.type;
        })
        
    }

    public ngOnDestroy(): void {
        
    }

}