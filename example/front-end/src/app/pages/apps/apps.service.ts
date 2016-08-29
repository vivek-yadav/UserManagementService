import {Injectable} from '@angular/core';

@Injectable()
export class AppsService {

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

    constructor() {
    }

    getApps() {
        return this.apps;
    }
}