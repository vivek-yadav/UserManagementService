import {Injectable} from '@angular/core';

@Injectable()
export class UsersService {

    private appsMap: any = {};
    private users = [
        {
            Username: "User1",
            Email: "user1@mathologic.com",
            TOC: "2016-01-01",
            TTL: "1y-3M-10d-4h-5m-6s",
            AppAccess: [
                { TOC: "2016-01-01", TTL: "1y-3M-10d-4h-5m-6s", Token: "CrewLogic:2345:2456", RoleId: "2", Paths: [{ Path: "crewlogic/api/admin/6/*", AccessLevel: 6 }] },
                { TOC: "2016-01-01", TTL: "1y-3M-10d-4h-5m-6s", Token: "LocoLogic:2345:2456", RoleId: "4", Paths: [{ Path: "locologic/api/admin/6/*", AccessLevel: 6 }] },
                { TOC: "2016-01-01", TTL: "1y-3M-10d-4h-5m-6s", Token: "TrainLogic:2345:2456", RoleId: "6", Paths: [{ Path: "trainlogic/api/admin/6/*", AccessLevel: 6 }] }
            ]
        },
        {
            Username: "User2",
            Email: "user2@mathologic.com",
            TOC: "2016-07-05",
            TTL: "1y-3M-10d-4h-5m-6s",
            AppAccess: [
                { TOC: "2016-07-05", TTL: "1y-3M-10d-4h-5m-6s", Token: "CrewLogic:2345:2456", RoleId: "1", Paths: [{ Path: "crewlogic/api/user/1/*", AccessLevel: 1 }] },
                { TOC: "2016-07-05", TTL: "1y-3M-10d-4h-5m-6s", Token: "LocoLogic:2345:2456", RoleId: "3", Paths: [{ Path: "locologic/api/user/2/*", AccessLevel: 2 }] },
                { TOC: "2016-07-05", TTL: "1y-3M-10d-4h-5m-6s", Token: "TrainLogic:2345:2456", RoleId: "5", Paths: [{ Path: "trainlogic/api/user/3/*", AccessLevel: 3 }] }
            ]
        },
        {
            Username: "User3",
            Email: "user3@mathologic.com",
            TOC: "2016-08-25",
            TTL: "1y-3M-10d-4h-5m-6s",
            AppAccess: [
                { TOC: "2016-08-25", TTL: "1y-3M-10d-4h-5m-6s", Token: "CrewLogic:2345:2456", RoleId: "1", Paths: [{ Path: "crewlogic/api/user/4/*", AccessLevel: 4 }] },
                { TOC: "2016-08-25", TTL: "1y-3M-10d-4h-5m-6s", Token: "LocoLogic:2345:2456", RoleId: "3", Paths: [{ Path: "locologic/api/user/5/*", AccessLevel: 5 }] },
                { TOC: "2016-08-25", TTL: "1y-3M-10d-4h-5m-6s", Token: "TrainLogic:2345:2456", RoleId: "5", Paths: [{ Path: "trainlogic/api/user/6/*", AccessLevel: 6 }] }
            ]
        }
    ];

    private roles = {
        CrewLogic: [{ id: 1, name: "user" },
            { id: 2, name: "admin" }],
        LocoLogic: [{ id: 3, name: "user" },
            { id: 4, name: "admin" }],
        TrainLogic: [{ id: 5, name: "user" },
            { id: 6, name: "admin" }]
    }

    private apps = [
        {
            Name: "CrewLogic",
            Description: "About Crew Logic",
            Token: "1248:7894:4657",
            TOC: "2016-01-01",
            TTL: "1y-3M-10d-4h-5m-6s",
            Roles: [
                {
                    _id: "1",
                    Name: "user",
                    Description: "About user role",
                    Paths: [{ Path: "crewlogic/api/user/0/*", AccessLevel: 0 },
                        { Path: "crewlogic/api/user/1/*", AccessLevel: 1 }, { Path: "crewlogic/api/user/2/*", AccessLevel: 2 },
                        { Path: "crewlogic/api/user/3/*", AccessLevel: 3 }, { Path: "crewlogic/api/user/4/*", AccessLevel: 4 },
                        { Path: "crewlogic/api/user/5/*", AccessLevel: 5 }, { Path: "crewlogic/api/user/6/*", AccessLevel: 6 }]
                },
                {
                    _id: "2",
                    Name: "admin",
                    Description: "About admin role",
                    Paths: [{ Path: "crewlogic/api/admin/6/*", AccessLevel: 6 }]
                }
            ]
        },
        {
            Name: "LocoLogic",
            Description: "About Loco Logic",
            Token: "1248:7894:4657",
            TOC: "2016-01-01",
            TTL: "1y-3M-10d-4h-5m-6s",
            Roles: [
                {
                    _id: "3",
                    Name: "user",
                    Description: "About user role",
                    Paths: [{ Path: "locologic/api/user/0/*", AccessLevel: 0 },
                        { Path: "locologic/api/user/1/*", AccessLevel: 1 }, { Path: "locologic/api/user/2/*", AccessLevel: 2 },
                        { Path: "locologic/api/user/3/*", AccessLevel: 3 }, { Path: "locologic/api/user/4/*", AccessLevel: 4 },
                        { Path: "locologic/api/user/5/*", AccessLevel: 5 }, { Path: "locologic/api/user/6/*", AccessLevel: 6 }]
                },
                {
                    _id: "4",
                    Name: "admin",
                    Description: "About admin role",
                    Paths: [{ Path: "locologic/api/admin/6/*", AccessLevel: 6 }]
                }
            ]
        },
        {
            Name: "TrainLogic",
            Description: "About Train Logic",
            Token: "1248:7894:4657",
            TOC: "2016-01-01",
            TTL: "1y-3M-10d-4h-5m-6s",
            Roles: [
                {
                    _id: "5",
                    Name: "user",
                    Description: "About user role",
                    Paths: [{ Path: "locologic/api/user/0/*", AccessLevel: 0 },
                        { Path: "locologic/api/user/1/*", AccessLevel: 1 }, { Path: "trainlogic/api/user/2/*", AccessLevel: 2 }, { Path: "trainlogic/api/user/1/*", AccessLevel: 2 },
                        { Path: "trainlogic/api/user/3/*", AccessLevel: 3 }, { Path: "trainlogic/api/user/4/*", AccessLevel: 4 },
                        { Path: "trainlogic/api/user/5/*", AccessLevel: 5 }, { Path: "trainlogic/api/user/6/*", AccessLevel: 6 }]
                },
                {
                    _id: "6",
                    Name: "admin",
                    Description: "About admin role",
                    Paths: [{ Path: "trainlogic/api/admin/6/*", AccessLevel: 6 }]
                }
            ]
        }
    ]

    createAppsMap() {
        var appsLength = this.apps.length, app, rolesLength, role;
        for (var i = 0; i < appsLength; i++) {
            app = this.apps[i];
            this.appsMap[app.Name] = JSON.parse(JSON.stringify(app));
            this.appsMap[app.Name].Roles = {};
            rolesLength = app.Roles.length;
            for (var j = 0; j < rolesLength; j++) {
                role = app.Roles[j];
                this.appsMap[app.Name].Roles[role._id] = role;
            }
        }
    }
    constructor() {
        this.createAppsMap();
    }

    getUsers() {
        return this.users;
    }

    getAppsMap() {
        return this.appsMap;
    }

    getRoles() {
        return this.roles;
    }
}