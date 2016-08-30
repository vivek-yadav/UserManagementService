import {Component, ViewEncapsulation} from '@angular/core';
import { Http, Response } from '@angular/http';

import { Router } from '@angular/router';
import { LoggedInGuard } from '../../services/guard/logged-in.guard';
import { AuthHttp } from '../../services/authHttp/authHttp.service';

import {BaCard} from '../../theme/components';


@Component({
  selector: 'dashboard',
  pipes: [],
  directives: [BaCard],
  encapsulation: ViewEncapsulation.None,
  providers:[AuthHttp],
  styles: [require('./dashboard.scss')],
  template: require('./dashboard.html')
})
export class Dashboard {
  private data:string;
  
  constructor(private router:Router,private loggedInGuard:LoggedInGuard, private authHttp:AuthHttp) {
  	if(!this.loggedInGuard.canActivate()){
  		this.router.navigate(['/login/false']);
  	}
  }

  ngOnInit(){
    // this.authHttp.get("http://localhost:7080/auth/api/v_0_1/user/vivek")
    // .map(res => res.text())
    // .subscribe(
    //   data => this.data = data,
    //   err => this.logError(err),
    //   () => console.log('Random Quote Complete')
    // );
  }

  logError(err:any){
    console.log("ERROR: ",err)
  }

}
