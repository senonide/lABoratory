import { NgModule } from '@angular/core';

import { CommonModule } from '@angular/common';  

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import { RendererComponent } from './renderer/renderer.component';
import { LandingView } from './components/landing/landing-view.component';

import { NgxChartsModule } from '@swimlane/ngx-charts';

import { HttpClientModule } from "@angular/common/http";

import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDividerModule } from '@angular/material/divider';
import {MatListModule} from '@angular/material/list';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatDialogModule} from '@angular/material/dialog';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {ClipboardModule} from '@angular/cdk/clipboard';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomeComponent } from './components/home/home.component';
import { AuthComponent } from './components/auth/auth.component';
import { ProfileComponent } from './components/profile/profile.component';
import { ExperimentList } from './components/profile/profile-components/experiment-list/experiment-list.component';
import { ExperimentContent } from './components/profile/profile-components/experiment-content/experiment-content.component';
import { CreateExperiment } from './components/profile/profile-components/create-experiment/create-experiment.component';
import { ExperimentDetails } from './components/profile/profile-components/experiment-details/experiment-details.component';
import { KeyDialog } from './components/profile/profile-components/key-dialog/key-dialog.component';
import { DeleteDialog } from './components/profile/profile-components/delete-dialog/delete-dialog.component';
import { UpdateExperiment } from './components/profile/profile-components/update-experiment/update-experiment.component';
import { AssignmentDialog } from './components/profile/profile-components/assignments-dialog/assignment-dialog.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    AuthComponent,
    RendererComponent,
    ProfileComponent,
    ExperimentList,
    ExperimentContent,
    CreateExperiment,
    UpdateExperiment,
    ExperimentDetails,
    LandingView,
    KeyDialog,
    AssignmentDialog,
    DeleteDialog
  ],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatListModule,
    CommonModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatInputModule,
    MatDividerModule,
    MatCardModule,
    MatGridListModule,
    MatDialogModule,
    ClipboardModule,
    MatButtonModule,
    MatToolbarModule,
    MatExpansionModule,
    MatProgressSpinnerModule,
    MatIconModule,
    NgxChartsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
