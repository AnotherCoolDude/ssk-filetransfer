import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable, Subject } from 'rxjs';
import { map } from 'rxjs/operators';
import { Basecampproject } from '../model/project';
import { Todoset } from '../model/todoset';
import { Todolist } from '../model/todolist';
import { Todo } from '../model/todo';

@Injectable({
  providedIn: 'root'
})
export class BasecampService {

  private notificationsource = new Subject<string>();

  constructor(private httpClient: HttpClient) { }

  notification(): Observable<string> {
    return this.notificationsource.asObservable();
  }

  valid(): Observable<boolean> {
    return this.httpClient.get<boolean>(environment.gateway + '/bc/valid').pipe(
      map((obj: any) => {
        return obj.tokenValid;
      })
    );
  }

  login(shortname: string): Observable<string> {
    return this.httpClient.get<string>(environment.gateway + '/bc/login', {params: {['shortname']: shortname}}).pipe(
      map((obj: any) => {
        console.log(obj);
        if (obj.redirectURL) {
          this.notificationsource.next(shortname);
          return obj.redirectURL;
        }
        if (obj.error) {
          return obj.error;
        }
        this.notificationsource.next(shortname);
        return '';
      })
    );
  }

  projects(shortname: string): Observable<Basecampproject[]> {
    return this.httpClient.get<Basecampproject[]>(environment.gateway + '/bc/projects', {params: {['shortname']: shortname}}).pipe(
      map((obj: any) => {
        const pp: Basecampproject[] = [];
        for (const d of obj) {
          pp.push(new Basecampproject(d));
        }
        return pp;
      })
    );
  }

  todoset(link: string, shortname: string): Observable<Todoset> {
    return this.httpClient.get<Todoset>(environment.gateway + '/bc/link', {params: {['shortname']: shortname, ['link']: link }});
  }

  todolists(link: string, shortname: string): Observable<Todolist[]> {
    return this.httpClient.get<Todolist[]>(environment.gateway + '/bc/link', {params: {['shortname']: shortname, ['link']: link }});
  }

  todos(link: string, shortname: string): Observable<Todo[]> {
    return this.httpClient.get<Todo[]>(environment.gateway + '/bc/link', {params: {['shortname']: shortname, ['link']: link }}).pipe(
      map(tt => tt)
    );
  }


}
