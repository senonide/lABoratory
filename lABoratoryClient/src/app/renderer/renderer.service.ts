import * as THREE from 'three';
import {ElementRef, Injectable, NgZone, OnDestroy } from '@angular/core';

@Injectable({providedIn: 'root'})
export class RendererService implements OnDestroy{
    private canvas!: HTMLCanvasElement;
    private renderer!: THREE.WebGLRenderer;
    private camera!: THREE.PerspectiveCamera;
    private scene!: THREE.Scene;
    private ambientLight!: THREE.AmbientLight;

    private cube!: THREE.Mesh;

    private circle!: THREE.Object3D;
    private skelet!: THREE.Object3D;
    private particle!: THREE.Object3D;

    private frameId: number = 0;


    public constructor(private ngZone: NgZone){}

    public ngOnDestroy(): void {
        if(this.frameId!=null){
            cancelAnimationFrame(this.frameId);
        }
    }

    public createScene(canvas: ElementRef<HTMLCanvasElement>): void {
        this.canvas = canvas.nativeElement;

        this.renderer = new THREE.WebGLRenderer({
            canvas: this.canvas,
            alpha: true,
            antialias: true,
        });
        this.renderer.setSize(window.innerWidth, window.innerHeight);
        this.renderer.setClearColor(0x000000, 0.0);

        this.scene = new THREE.Scene();
       
        this.camera = new THREE.PerspectiveCamera(
            75, window.innerWidth/window.innerHeight, 0.1, 1000
        );

        this.camera.position.z = 400;
        this.scene.add(this.camera);

        this.circle = new THREE.Object3D();
        this.skelet = new THREE.Object3D();
        this.particle = new THREE.Object3D();

        this.circle.position.set(0, 0, 0);
        this.skelet.position.set(0, 0, 0);

        this.scene.add(this.circle);
        this.scene.add(this.skelet);
        this.scene.add(this.particle);

        var geometry = new THREE.TetrahedronGeometry(2, 0);
        var geom = new THREE.IcosahedronGeometry(7);
        var geom2 = new THREE.IcosahedronGeometry(14);

        var material = new THREE.MeshStandardMaterial({
            color: 0xffffff,
            //shading: THREE.FlatShading
        });

        for (var i = 0; i < 1000; i++) {
            var mesh = new THREE.Mesh(geometry, material);
            mesh.position.set(Math.random() - 0.5, Math.random() - 0.5, Math.random() - 0.5).normalize();
            mesh.position.multiplyScalar(90 + (Math.random() * 700));
            mesh.rotation.set(Math.random() * 2, Math.random() * 2, Math.random() * 2);
            this.particle.add(mesh);
        }

        var mat = new THREE.MeshStandardMaterial({
            color: 0xffffff,
            
        });

        var mat2 = new THREE.MeshStandardMaterial({
            color: 0xffffff,
            wireframe: true,
            side: THREE.DoubleSide
        });

        var planet = new THREE.Mesh(geom, mat);
        planet.scale.x = planet.scale.y = planet.scale.z = 16;
        this.circle.add(planet);

        var planet2 = new THREE.Mesh(geom2, mat2);
        planet2.scale.x = planet2.scale.y = planet2.scale.z = 10;
        this.skelet.add(planet2);

        var ambientLight = new THREE.AmbientLight(0x999999, 0.5 );
        this.scene.add(ambientLight);

        var lights = [];
        lights[0] = new THREE.DirectionalLight( 0xffffff, 0.5 );
        lights[0].position.set( 1, 0, 0 );
        lights[1] = new THREE.DirectionalLight( 0x11E8BB, 1 );
        lights[1].position.set( 0.75, 1, 0.5 );
        lights[2] = new THREE.DirectionalLight( 0x8200C9, 1 );
        lights[2].position.set( -0.75, -1, 0.5 );
        this.scene.add( lights[0] );
        this.scene.add( lights[1] );
        this.scene.add( lights[2] );

        
    }

    public animate(): void {
        this.ngZone.runOutsideAngular(() => {

            if(document.readyState !== 'loading') {
                this.render();
            } else {
                window.addEventListener('DOMContentLoaded', () => {
                    this.render();
                });
                window.addEventListener('resize', () => {
                    this.camera.aspect = window.innerWidth / window.innerHeight;
                    this.camera.updateProjectionMatrix();
                    this.renderer.setSize(window.innerWidth, window.innerHeight);
                    this.render();
                });
            }
        });
    }

    public render(): void {
        this.frameId = requestAnimationFrame(() => {
            this.render();
        });
        this.particle.rotation.x += 0.0000;
        this.particle.rotation.y -= 0.0040;
        this.circle.rotation.x -= 0.01;
        this.circle.rotation.y -= 0.01;
        this.skelet.rotation.x -= 0.005;
        this.skelet.rotation.y += 0.005;
        this.renderer.clear();

        this.renderer.render(this.scene, this.camera);
    }
}