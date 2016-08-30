import {Component, ViewEncapsulation} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import { Router,ActivatedRoute } from '@angular/router';

import { UserService } from '../../services/user/user.service';
import { LoggedInGuard } from '../../services/guard/logged-in.guard';

@Component({
  selector: 'forgot-pass',
  encapsulation: ViewEncapsulation.None,
  directives: [],
  styles: [require('./forgot-pass.scss')],
  template: require('./forgot-pass.html'),
})
export class ForgotPass {

  public form:FormGroup;
  public email:AbstractControl;
  public mobileNo:AbstractControl;
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
      'mobileNo': ['', Validators.compose([Validators.required, Validators.minLength(10)])]
    });

    this.email = this.form.controls['email'];
    this.mobileNo = this.form.controls['mobileNo'];
  }

  public onSubmit(values:Object):void {
    this.submitted = true;
    if (this.form.valid) {
      this.router.navigate(['/login']);
      // your code goes here
      // console.log(values);
      // this.userService.login(values["email"], values["mobileNo"]).subscribe((result) => {
      //   if (result) {
      //     this.router.navigate(['/login']);
      //   }
      // });
    }
  }
}
