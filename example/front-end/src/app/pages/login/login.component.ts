import {Component, ViewEncapsulation} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import { Router,ActivatedRoute } from '@angular/router';

import { UserService } from '../../services/user/user.service';
import { LoggedInGuard } from '../../services/guard/logged-in.guard';

@Component({
  selector: 'login',
  encapsulation: ViewEncapsulation.None,
  directives: [],
  styles: [require('./login.scss')],
  template: require('./login.html'),
})
export class Login {

  public form:FormGroup;
  public email:AbstractControl;
  public password:AbstractControl;
  public submitted:boolean = false;

  constructor(private userService: UserService,fb:FormBuilder, private router: Router,private loggedInGuard:LoggedInGuard,private aRoute: ActivatedRoute) {
    if(this.aRoute.snapshot.params["doLogOut"] && this.loggedInGuard.canActivate()){
      this.userService.logout()
    }
    if(this.loggedInGuard.canActivate()){
      this.router.navigate(['/pages/dashboard']);
    }
    this.form = fb.group({
      'email': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
      'password': ['', Validators.compose([Validators.required, Validators.minLength(4)])]
    });

    this.email = this.form.controls['email'];
    this.password = this.form.controls['password'];
  }

  public onSubmit(values:Object):void {
    this.submitted = true;
    if (this.form.valid) {
      // your code goes here
      console.log(values);
      this.userService.login(values["email"], values["password"]).subscribe((result) => {
        if (result) {
          this.router.navigate(['/pages/dashboard']);
        }
      });
    }
  }
}
