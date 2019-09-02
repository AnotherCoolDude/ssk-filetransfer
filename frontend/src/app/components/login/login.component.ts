import { Component, OnInit } from '@angular/core';
import { ProadService } from 'src/app/services/proad.service';
import { Employee } from 'src/app/model/employee';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  shortname: string;
  employeeList: Employee[];

  constructor(
    private proadService: ProadService,
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
        this.router.navigate(['/filetransfertable/' + ee[0].urno])
          .then((result: boolean) => console.log(result));
      }
    }
  }

}
