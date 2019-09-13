import { Bucket, Parent, Creator } from './general';

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
  bookmark_url: string;
  position: number;
  bucket: Bucket;
  creator: Creator;
  completed: boolean;
  completed_ratio: string;
  name: string;
  todolists_count: number;
  todolists_url: string;
  app_todoslists_url: string;
}
