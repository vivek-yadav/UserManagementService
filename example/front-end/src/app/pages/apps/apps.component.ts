import {Component, ViewEncapsulation, ChangeDetectorRef} from '@angular/core';
import { Http, Response } from '@angular/http';
import { Router } from '@angular/router';

import {InlineEditComponent} from '../../utils/UI/inline-edit/inline-edit.component.ts';

import { LoggedInGuard } from '../../services/guard/logged-in.guard';
import { AuthHttp } from '../../services/authHttp/authHttp.service';
import { AppsService } from './apps.service';

import {BaCard} from '../../theme/components';

@Component({
  selector: 'dashboard',
  pipes: [],
  directives: [BaCard, InlineEditComponent],
  encapsulation: ViewEncapsulation.None,
  providers: [AuthHttp, AppsService],
  styles: [require('./apps.scss')],
  template: require('./apps.html')
})
export class Apps {
  private data: string;
  appsService: AppsService;
  apps: any; viewPathsAddress: any = {};
  cdr: ChangeDetectorRef;
  viewPaths: any = [];

  constructor(private router: Router, private loggedInGuard: LoggedInGuard, private authHttp: AuthHttp,
    private as: AppsService, private CDR: ChangeDetectorRef) {
    this.appsService = as; this.cdr = CDR;
    if (!this.loggedInGuard.canActivate()) {
      this.router.navigate(['/login/false']);
    }
  }

  ngAfterViewInit() {
    this.apps = this.appsService.getApps();
    this.cdr.detectChanges();
  }

  logError(err: any) {
    console.log("ERROR: ", err)
  }

  /**
   * Function to update app details on edit
   */
  updateApp(app, value, keyL1, keyL2, keyL3, keyL4, keyL5) {
    if (keyL5 != undefined) {
      app[keyL1][keyL2][keyL3][keyL4][keyL5] = value;
    } else if (keyL4 != undefined) {
      app[keyL1][keyL2][keyL3][keyL4] = value;
    } else if (keyL3 != undefined) {
      app[keyL1][keyL2][keyL3] = value;
    } else if (keyL2 != undefined) {
      app[keyL1][keyL2] = value;
    } else {
      app[keyL1] = value;
    }
  }

  /**
   * Function to set paths to view in modal
   */
  setModalPaths(app, j, paths) {
    this.viewPaths = paths;
    this.viewPathsAddress = { app: app, roleIndex: j };
  }

  /**
   * Function to addd new role to app
   */
  addNewRole(app) {
    app.Roles.push({
      Name: "",
      Description: "",
      Paths: [{ Path: "", AccessLevel: "", isNew: true }],
      isNew: true
    })
  }

  /**
   * Function to delete role from app
   */
  deleteRole(app, j) {
    app.Roles.splice(j, 1);
  }

  /**
   * Function to add new path to a role in app
   */
  addNewPath() {
    var app = this.viewPathsAddress.app;
    app.Roles[this.viewPathsAddress.roleIndex].Paths.push({ Path: "", AccessLevel: "", isNew: true });
  }
  /**
   * Function to delete pathe from role in app
   */
  deletePath(app, roleIndex, pathIndex) {
    app.Roles[roleIndex].Paths.splice(pathIndex, 1);
  }

  /**
   * Function to add new app
   */
  addNewApp() {
    var date = new Date();
    this.apps.push({
      isNew : true,
      Name: "",
      Description: "",
      Token: "",
      TOC: date.getFullYear()+"-"+date.getMonth()+"-"+date.getDate(),
      TTL: "",
      Roles: [
        {
          isNew : true,
          Name: "",
          Description: "",
          Paths: [{ Path: "", AccessLevel: "", isNew: true }]
        }
      ]
    })
  }

  /**
   * Function to delete pathe from role in app
   */
  deleteApp(appIndex) {
    this.apps.splice(appIndex, 1);
  }
}
