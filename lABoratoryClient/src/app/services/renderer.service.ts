import * as THREE from 'three';
import {ElementRef, Injectable, NgZone, OnDestroy } from '@angular/core';

@Injectable({providedIn: 'root'})
export class RendererService implements OnDestroy{
    private canvas!: HTMLCanvasElement;
    private renderer!: THREE.WebGLRenderer;
    private camera!: THREE.PerspectiveCamera;
    private scene!: THREE.Scene;

    private cube!: THREE.Object3D;
    private cubeMesh!: THREE.Object3D;
    private particles!: THREE.Object3D;

    private frameId: number = 0;

    private isRendering: boolean = false;


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

        this.cube = new THREE.Object3D();
        this.cubeMesh = new THREE.Object3D();
        this.particles = new THREE.Object3D();

        this.cube.position.set(0, 20, 0);
        this.cubeMesh.position.set(0, 20, 0);
        this.particles.position.set(0, 20, 0);

        this.scene.add(this.cube);
        this.scene.add(this.cubeMesh);
        this.scene.add(this.particles);

        var particleShape = new THREE.BoxGeometry(6, 6, 6);
        var mainCubeGeometry = new THREE.IcosahedronGeometry(7);
        var meshCubeGeometry = new THREE.IcosahedronGeometry(14);

        var material = new THREE.MeshStandardMaterial({
            color: 0xffffff,
        });

        for (var i = 0; i < 1200; i++) {
            var mesh = new THREE.Mesh(particleShape, material);
            mesh.position.set((Math.random() - 0.5), (Math.random() - 0.5) * 0.5, (Math.random() - 0.5)).normalize();
            mesh.position.multiplyScalar(200 + (Math.random() * 700));
            mesh.rotation.set(Math.random() * 2, Math.random() * 2, Math.random() * 2);
            this.particles.add(mesh);
        }

        var mat = new THREE.MeshStandardMaterial({
            color: 0xffffff,
            
        });

        var mat2 = new THREE.MeshStandardMaterial({
            color: 0xffffff,
            wireframe: true,
            side: THREE.DoubleSide
        });

        var planet = new THREE.Mesh(mainCubeGeometry, mat);
        planet.scale.x = planet.scale.y = planet.scale.z = 16;
        this.cube.add(planet);

        var planet2 = new THREE.Mesh(meshCubeGeometry, mat2);
        planet2.scale.x = planet2.scale.y = planet2.scale.z = 10;
        this.cubeMesh.add(planet2);

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
                if(!this.isRendering){
                    this.render();
                    this.isRendering = true;
                }
            } else {
                window.addEventListener('DOMContentLoaded', () => {
                    if(!this.isRendering){
                        this.render();
                        this.isRendering = true;
                    }
                });
                
               
            }

            window.addEventListener('resize', () => {
                this.camera.aspect = window.innerWidth / window.innerHeight;
                this.camera.updateProjectionMatrix();

                this.renderer.setSize( window.innerWidth, window.innerHeight );
                console.log("resize");
            })
        });
    }

    public render(): void {
        this.frameId = requestAnimationFrame(() => {
            this.render();
        });
        this.particles.rotation.x += 0.0000;
        this.particles.rotation.y -= 0.0020;
        this.cube.rotation.x -= 0.01;
        this.cube.rotation.y -= 0.01;
        this.cubeMesh.rotation.x -= 0.003;
        this.cubeMesh.rotation.y += 0.003;
        this.renderer.clear();

        this.renderer.render(this.scene, this.camera);
    }
}