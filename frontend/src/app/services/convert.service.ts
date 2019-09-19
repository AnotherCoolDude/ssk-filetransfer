import { Injectable } from '@angular/core';
import { Todo } from '../model/todo';
import { Task } from '../model/task';
import { ProadService, StatusCode } from './proad.service';
import { Employee } from '../model/employee';
import { Project } from '../model/project';

@Injectable({
  providedIn: 'root'
})
export class ConvertService {

  employees: Employee[];

  constructor(private proad: ProadService) {
    proad.getEmployees().subscribe(ee => this.employees = ee);
   }

  task(todo: Todo, project: Project, urno?: number): Task {
    const assignee = this.employees.filter(e => e.firstname + ' ' + e.firstname === todo.assignees[0]);
    const t: Task = {
      urno: urno || 0,
      urno_manager: this.proad.currentEmployee.urno,
      urno_company: 0,
      urno_project: project.urno,
      urno_service_code: 0,
      urno_responsible: assignee[0].urno,
      shortinfo: todo.title,
      from_datetime: todo.created_at,
      until_datetime: todo.due_on || new Date().toISOString(),
      reminder_datetime: '',
      status: String(StatusCode.vorbereitung),
      priority: '',
      description: todo.description,
    }
    return t;
  }

  // todo(task: Task): Todo {
  //   cont t: Todo = {

  //   }
  // }

}
