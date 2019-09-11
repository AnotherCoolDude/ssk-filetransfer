import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Basecampproject } from '../model/project';

@Injectable({
  providedIn: 'root'
})
export class BasecampService {

  constructor(private httpClient: HttpClient) { }

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
          return obj.redirectURL;
        }
        if (obj.error) {
          return obj.error;
        }
        return '';
      })
    );
  }

  projects(): Observable<Basecampproject[]> {
    return this.httpClient.get<Basecampproject[]>(environment.gateway + '/bc/projects').pipe(
      map((obj: any) => {
        const pp: Basecampproject[] = [];
        for (const d of obj) {
          pp.push(new Basecampproject(d));
        }
        return pp;
      })
    );
  }



}
