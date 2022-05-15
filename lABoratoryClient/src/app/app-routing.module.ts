import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthComponent } from './components/auth/auth.component';
import { HomeComponent } from './components/home/home.component';
import { ProfileComponent } from './components/profile/profile.component';
import { WebExample } from './components/web-example/web-example.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'example',
    component: WebExample
  },
  {
    path: 'auth/:type',
    component: AuthComponent
  },
  {
    path: 'profile',
    component: ProfileComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
