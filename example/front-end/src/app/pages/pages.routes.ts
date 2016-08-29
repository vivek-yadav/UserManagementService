import {RouterConfig} from '@angular/router';
import {Dashboard} from './dashboard/dashboard.component';
import {Users} from './users/users.component';
import {Apps} from './apps/apps.component';
import {Pages} from './pages.component';

import { LoggedInGuard } from '../services/guard/logged-in.guard';

//noinspection TypeScriptValidateTypes
export const PagesRoutes:RouterConfig = [
  {
    path: 'pages',
    component: Pages,
    canActivate: [LoggedInGuard],
    children: [
      {
        path: 'dashboard',
        component: Dashboard,
        data: {
          menu: {
            title: 'Dashboard',
            icon: 'ion-android-home',
            selected: false,
            expanded: false,
            order: 0
          }
        }
      },{
        path: 'users',
        component: Users,
        data: {
          menu: {
            title: 'Users',
            icon: 'fa fa-user',
            selected: false,
            expanded: false,
            order: 0
          }
        }
      },{
        path: 'apps',
        component: Apps,
        data: {
          menu: {
            title: 'Apps',
            icon: 'fa fa-cloud',
            selected: false,
            expanded: false,
            order: 0
          }
        }
      }
      // {
      //   path: 'editors',
      //   component: Editors,
      //   data: {
      //     menu: {
      //       title: 'Editors',
      //       icon: 'ion-edit',
      //       selected: false,
      //       expanded: false,
      //       order: 100,
      //     }
      //   },
      //   children: [
      //     {
      //       path: 'ckeditor',
      //       component: Ckeditor,
      //       data: {
      //         menu: {
      //           title: 'CKEditor',
      //         }
      //       }
      //     }
      //   ]
      // },
    ]
  }
];
