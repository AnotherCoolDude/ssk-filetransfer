import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { ProadService } from 'src/app/services/proad.service';
import { switchMap } from 'rxjs/operators';
import { Project } from 'src/app/model/projects';
import { Observable } from 'rxjs';
import { Employee } from 'src/app/model/employee';

@Component({
  selector: 'app-filetransfertable',
  templateUrl: './filetransfertable.component.html',
  styleUrls: ['./filetransfertable.component.css']
})
export class FiletransfertableComponent implements OnInit {

  projects$: Observable<Project[]>;
  employees$: Observable<Employee[]>;
  employeeUrno: number;

  displayedColumns: string[] = [ 'urno', 'title', 'status' ];
  dataSource = this.projects$;

  constructor(
    private route: ActivatedRoute,
    private proadService: ProadService
  ) { }

  ngOnInit() {
    this.projects$ = this.route.paramMap.pipe(
      switchMap((params: ParamMap) => {
        this.employeeUrno = +params.get('urno');
        return this.proadService.getProjectsForEmployee( +params.get('urno'));
      })
    );
    this.employees$ = this.proadService.getEmployees();
  }

}