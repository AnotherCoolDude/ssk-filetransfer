import { Todoset } from './todoset';
import { Todo } from './todo';

export class Project {
    urno: number;
    status: string;
    // tslint:disable-next-line:variable-name
    project_name: string;
    // tslint:disable-next-line:variable-name
    order_date: Date;
    projectno: string;
    description: string;
    content: string[];
}

export class Basecampproject {
    id: number;
    status: string;
    createdAt: Date;
    updatedAt: Date;
    name: string;
    description: string;
    purpose: string;
    bookmarkURL: URL;
    url: URL;
    appURL: URL;
    dock: {
        id: number;
        title: string;
        name: string;
        enabled: boolean;
        position: number;
        url: URL;
        app_url: URL;
    };
    todoset: Todoset;
    todos: Todo[];

    constructor(obj: any) {
        this.id = obj.id;
        this.status = obj.status;
        this.createdAt = obj.created_at;
        this.updatedAt = obj.updated_at;
        this.name = obj.name;
        this.description = obj.description;
        this.purpose = obj.purpose;
        this.bookmarkURL = obj.bookmark_url;
        this.url = obj.url;
        this.appURL = obj.app_url;
        for (const d of obj.dock) {
            if (d.name === 'todoset') {
                this.dock = d;
            }
        }
        this.todos = [];
    }
}
