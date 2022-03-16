import {Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import {RendererService} from '../services/renderer.service';

@Component({
    selector: 'app-renderer',
    templateUrl: './renderer.component.html',
    styleUrls: ['./renderer.component.css']
})
export class RendererComponent implements OnInit {

    @ViewChild('rendererCanvas', { static: true })
    public rendererCanvas!: ElementRef<HTMLCanvasElement>;

    public constructor(private rendServ: RendererService){}

    public ngOnInit(): void {
        this.rendServ.createScene(this.rendererCanvas);
        this.rendServ.animate();
    }

}