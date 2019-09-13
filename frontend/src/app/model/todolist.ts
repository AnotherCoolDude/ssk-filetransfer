import { Bucket, Parent, Creator } from './general';

export interface Todolist {
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
  completed_ratio: string;
  name: string;
  todos_url: string;
  groups_url: string;
  app_todos_url: string;
}
