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
