import {Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';

@Component({
    selector: 'landing-view',
    templateUrl: './landing-view.component.html',
    styleUrls: ['./landing-view.component.css']
})
export class LandingView implements OnInit, OnDestroy {

    @ViewChild('rendererCanvas', { static: true })
    public rendererCanvas!: ElementRef<HTMLCanvasElement>;

    public constructor(){}

    public ngOnInit(): void {
        
    }

    public ngOnDestroy(): void {
        
    }

}