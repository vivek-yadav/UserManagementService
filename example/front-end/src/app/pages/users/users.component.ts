import {Component, ViewEncapsulation, ChangeDetectorRef} from '@angular/core';
import { Http, Response } from '@angular/http';
import { Router } from '@angular/router';

import {InlineEditComponent} from '../../utils/UI/inline-edit/inline-edit.component.ts';

import { LoggedInGuard } from '../../services/guard/logged-in.guard';
import { AuthHttp } from '../../services/authHttp/authHttp.service';
import { UsersService } from './users.service';

import {BaCard} from '../../theme/components';

@Component({
  selector: 'dashboard',
  pipes: [],
  directives: [BaCard, InlineEditComponent],
  encapsulation: ViewEncapsulation.None,
  providers: [AuthHttp, UsersService],
  styles: [require('./users.scss')],
  template: require('./users.html')
})
export class Users {
  private data: string;
  usersService: UsersService;
  users: any; appsMap: any; roles: any;
  cdr: ChangeDetectorRef;
  viewPaths: any = [];

  constructor(private router: Router, private loggedInGuard: LoggedInGuard, private authHttp: AuthHttp,
    private us: UsersService, private CDR: ChangeDetectorRef) {
    this.usersService = us; this.cdr = CDR;
    if (!this.loggedInGuard.canActivate()) {
      this.router.navigate(['/login/false']);
    }
  }

  ngAfterViewInit() {
    this.users = this.usersService.getUsers();
    this.appsMap = this.usersService.getAppsMap();
    this.roles = this.usersService.getRoles();
    this.cdr.detectChanges();
  }

  logError(err: any) {
    console.log("ERROR: ", err)
  }

  /**
   * Function to return list of roles by appname
   */
  getRolesByAppName(appName) {
    return this.roles[appName];
  }

  getAppName(token) {
    return token.split(':')[0];
  }

  /**
   * Function to update paths on role change
   */
  onRoleUpdate(user, index, appName, role) {
    user.AppAccess[index].Paths = this.appsMap[appName].Roles[role].Paths;
  }

  /**
   * Function to update user details on edit
   */
  updateUser(user, value, keyL1, keyL2, keyL3) {
    if (keyL3 != undefined) {
      user[keyL1][keyL2][keyL3] = value;
    } else if (keyL2 != undefined) {
      user[keyL1][keyL2] = value;
    } else {
      user[keyL1] = value;
    }
  }

  /**
   * Function to set paths to view in modal
   */
  setModalPaths(paths) {
    this.viewPaths = paths;
  }
}
