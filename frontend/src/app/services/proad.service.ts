import { Injectable } from '@angular/core';

import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Employee } from '../model/employee';
import { environment } from 'src/environments/environment';
import { Project } from '../model/project';

@Injectable({
  providedIn: 'root'
})
export class ProadService {

  private employeedataChanged = true;
  private projectdataChanged = true;
  private cachedEmployees: Employee[];
  private cachedProjects: Project[];

  constructor(private httpClient: HttpClient) { }

  dataHasChanged() {
    this.employeedataChanged = true;
    this.projectdataChanged = true;
  }

  getEmployees(): Observable<Employee[]> {
    if (!this.employeedataChanged) {
      return new Observable((obs) => {
        obs.next(this.cachedEmployees);
        obs.complete();
      });
    }
    return this.httpClient.get(environment.gateway + '/proad/employees').pipe(
      map((personList: any) => {
        const pp: Employee[] = [];
        for (const p of personList.person_list) {
          pp.push(p);
        }
        this.cachedEmployees = pp;
        this.employeedataChanged = false;
        return pp;
      })
    );
  }

  getProjectsForEmployee(urno: number): Observable<Project[]> {
    if (!this.projectdataChanged) {
      return new Observable((obs) => {
        console.log(this.cachedProjects);
        obs.next(this.cachedProjects);
        obs.complete();
      });
    }
    return this.httpClient.get(environment.gateway + '/proad/projects/' + urno).pipe(
      map((projectList: any) => {
        const pp: Project[] = [];
        for (const p of projectList.project_list) {
          pp.push(p);
        }
        this.cachedProjects = pp;
        this.projectdataChanged = false;
        return pp;
      })
    );
  }

}
