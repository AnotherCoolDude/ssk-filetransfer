import { Injectable } from '@angular/core';
import { Todo } from '../model/todo';
import { Task } from '../model/task';

@Injectable({
  providedIn: 'root'
})
export class ConvertService {

  constructor() { }

  task(todo: Todo, urno?: number): Task {
    const t: Task = {
      urno: urno || 0,
      urno_manager: 

    }
  }

  todo(task: Task): Todo {

  }

}
