import {provideRouter, RouterConfig} from '@angular/router';
import {LoginRoutes} from "./pages/login/login.routes";
import {PagesRoutes} from "./pages/pages.routes";
import {RegisterRoutes} from "./pages/register/register.routes";
import {ForgotPassRoutes} from "./pages/forgot-pass/forgot-pass.routes";

export const routes:RouterConfig = [
  ...LoginRoutes,
  ...RegisterRoutes,
  ...PagesRoutes,
  ...ForgotPassRoutes,
  {
    path: '**',
    redirectTo: '/login/false'
  },
];

export const APP_ROUTER_PROVIDERS = [
  provideRouter(routes)
];
