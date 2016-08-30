import {provide,Injectable} from '@angular/core';
import {HTTP_PROVIDERS, Http, Request, RequestOptionsArgs, Response, XHRBackend, RequestOptions, ConnectionBackend, Headers} from '@angular/http';
import {Router} from '@angular/router';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class AuthHttp {
     
    constructor(private http:Http,private _router:Router) {
    }
 
    request(url: string | Request, options?: RequestOptionsArgs): Observable<Response> {
        return this.intercept(this.http.request(url, options));
    }
 
    get(url: string, options?: RequestOptionsArgs): Observable<Response> {
        return this.intercept(this.http.get(url,this.getRequestOptionArgs(options)));
    }
 
    post(url: string, body: string, options?: RequestOptionsArgs): Observable<Response> {   
        return this.intercept(this.http.post(url, body, this.getRequestOptionArgs(options)));
    }
 
    put(url: string, body: string, options?: RequestOptionsArgs): Observable<Response> {
        return this.intercept(this.http.put(url, body, this.getRequestOptionArgs(options)));
    }
 
    delete(url: string, options?: RequestOptionsArgs): Observable<Response> {
        return this.intercept(this.http.delete(url, options));
    }
    
    getRequestOptionArgs(options?: RequestOptionsArgs) : RequestOptionsArgs {
        if (options == null) {
            options = new RequestOptions();
        }
        if (options.headers == null) {
            options.headers = new Headers();
        }
        options.headers.append('Content-Type', 'application/json');
        this.createAuthorizationHeader(options.headers);
        return options;
    }

    createAuthorizationHeader(headers:Headers) {
        if(!!localStorage.getItem('auth_token')){
            headers.append('Authorization', 'Bearer ' + localStorage.getItem('auth_token'));   
        }
    }
 
    intercept(observable: Observable<Response>): Observable<Response> {
        return observable.catch((err, source) => {
            if (err.status  == 401) {
                this._router.navigate(['/login/true']);
                return Observable.empty();
            } else {
                return Observable.throw(err);
            }
        });
 
    }
}