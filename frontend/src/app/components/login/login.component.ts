import { Component, OnInit } from '@angular/core';
import { ProadService, StatusCode } from 'src/app/services/proad.service';
import { Employee } from 'src/app/model/employee';
import { Router } from '@angular/router';
import { BasecampService } from 'src/app/services/basecamp.service';
import { Subject, Observable } from 'rxjs';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  shortname: string;
  employeeList: Employee[];

  loginbctext = 'Login to Basecamp';
  valid = false;

  private notificationSource = new Subject<string>();

  notification(): Observable<string> {
    return this.notificationSource.asObservable();
  }

  constructor(
    private proadService: ProadService,
    private basecampservice: BasecampService,
    private router: Router
  ) { }

  ngOnInit() {
    this.proadService.getEmployees().subscribe(eList => {
      this.employeeList = eList;
      console.log(this.employeeList);
    });
  }

  checkShortname() {
    if (this.shortname) {
      const ee = this.employeeList.filter(e => e.shortname === this.shortname);
      if (ee.length === 1) {
        console.log(ee[0]);
        this.proadService.dataHasChanged();
        this.router.navigate(['/filetransfertable/' + ee[0].urno]);
      }
    }
  }

  logintobc() {
    if (this.shortname) {
      const ee = this.employeeList.filter(e => e.shortname === this.shortname);
      if (ee.length === 1) {
        this.basecampservice.login(ee[0].shortname).subscribe(result => {
          if (result === '') {
            this.loginbctext = 'go to table';
            this.router.navigate(['/basecamptable/' + this.shortname]);
            console.log(result);
          }
          if (result.startsWith('http')) {
            window.open(result);
            return;
          } else {
            console.log('shortname missing');
          }
        });
      }
    }
  }
}

