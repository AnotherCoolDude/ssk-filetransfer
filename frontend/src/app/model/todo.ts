import { Parent, Bucket, Creator } from './general';
import { Task } from './task';

export interface Todo {
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
