import { Injectable } from '@angular/core';

import { HttpClient, HttpParams } from '@angular/common/http';
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

  getFilteredProjects(code: StatusCode, startDate: Date, endDate: Date): Observable<Project[]> {
    if (!this.projectdataChanged) {
      return new Observable((obs) => {
        console.log(this.cachedProjects);
        obs.next(this.cachedProjects);
        obs.complete();
      });
    }
    let filterParams = new HttpParams();
    filterParams = filterParams.append('status', String(code));
    filterParams = filterParams.append('startDate', this.ISODateString(startDate));
    filterParams = filterParams.append('endDate', this.ISODateString(endDate));
    return this.httpClient.get(environment.gateway + '/proad/projects', { params: filterParams }).pipe(
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

  getProjectByProjectnr(projectnr: string): Observable<Project> {
    let prnrParam = new HttpParams();
    prnrParam = prnrParam.append('projectnr', projectnr);
    return this.httpClient.get<Project>(environment.gateway + '/proad/project', {params: prnrParam}).pipe(
      map((projectlist: any) => {
        if (projectlist.error) {
          console.log(projectlist.error);
          return new Project();
        }
        return projectlist.project_list[0];
      })
    );
  }

  getContentForProject(projectnr: string): Observable<string[]> {
    let prnrParam = new HttpParams();
    prnrParam = prnrParam.append('projectnr', projectnr);
    return this.httpClient.get<string[]>(environment.gateway + '/files/project', {params: prnrParam});
  }

  private ISODateString(d: Date) {
    function pad(n: number) { return n < 10 ? '0' + n : n; }
    return d.getUTCFullYear() + '-'
      + pad(d.getUTCMonth() + 1) + '-'
      + pad(d.getUTCDate()) + 'T'
      + pad(d.getUTCHours()) + ':'
      + pad(d.getUTCMinutes()) + ':'
      + pad(d.getUTCSeconds()) + 'Z';
  }

}

export enum StatusCode {
  // StatusNone does not filter any StatusCode
  none = 0,
  // StatusVorbereitung 100
  vorbereitung = 100,
  // StatusAngebot 200
  angebot = 200,
  // StatusDurchführung 300
  durchführung = 300,
  // StatusDurchführungTE Durchführung Teilabrechnung erfolgt 301
  durchführungTE = 301,
  // StatusGeliefert 400
  geliefert = 400,
  // StatusAbgerechnet 500
  abgerechnet = 500,
  // StatusAbgerechnetVP Abgerechnet - Verrechnung mit anderem Project 501
  abgerechnetVP = 501,
  // StatusAbgebrochen 600
  abgebrochen = 600,
  // StatusAbgebrochenNR Abgebrochen - nicht realisiert 601
  abgebrochenNR = 601,
  // StatusAbgebrochenNW Abgebrochen - nicht weiterberechenbar 602
  abgebrochenNW = 602,
}
