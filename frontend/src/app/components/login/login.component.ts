import { Component, OnInit } from '@angular/core';
import { ProadService, StatusCode } from 'src/app/services/proad.service';
import { Employee } from 'src/app/model/employee';
import { Router } from '@angular/router';
import { BasecampService } from 'src/app/services/basecamp.service';

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

  constructor(
    private proadService: ProadService,
    private basecampservice: BasecampService,
    private router: Router
  ) { }

  ngOnInit() {
    this.basecampservice.valid().subscribe(valid => {
      if (valid) {
        this.loginbctext = 'go to table';
      } else {
        this.loginbctext = 'login';
      }
    });

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
        this.router.navigate(['/filetransfertable/' + ee[0].urno])
          .then((result: boolean) => console.log(result));
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
            this.router.navigate(['/basecamptable']);
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

