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
    todosets: Todoset[];

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
    }
}

export interface Todoset {
  id: number;
  status: string;
  created_at: string;
  updated_at: string;
  title: string;
  inherits_status: boolean;
  type: string;
  url: string;
  app_url: string;
  comments_count: number;
  comments_url: string;
  position: number;
  parent: Parent;
  bucket: Bucket;
  creator: Creator;
  description: string;
  bookmark_url: string;
  subscription_url: string;
  completed: boolean;
  content: string;
  starts_on?: any;
  due_on?: any;
  assignees: any[];
  completion_subscribers: any[];
  completion_url: string;
}

export interface Creator {
  id: number;
  attachable_sgid: string;
  name: string;
  email_address: string;
  personable_type: string;
  title: string;
  bio: string;
  created_at: string;
  updated_at: string;
  admin: boolean;
  owner: boolean;
  time_zone: string;
  avatar_url: string;
  company: Company;
}

export interface Company {
  id: number;
  name: string;
}

export interface Bucket {
  id: number;
  name: string;
  type: string;
}

export interface Parent {
  id: number;
  title: string;
  type: string;
  url: string;
  app_url: string;
}
