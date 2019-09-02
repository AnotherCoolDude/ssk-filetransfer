import { Component, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { ProadService } from 'src/app/services/proad.service';
import { switchMap } from 'rxjs/operators';
import { Project } from 'src/app/model/project';
import { Observable } from 'rxjs';
import { Employee } from 'src/app/model/employee';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { throwToolbarMixedModesError } from '@angular/material/toolbar';

@Component({
  selector: 'app-filetransfertable',
  templateUrl: './filetransfertable.component.html',
  styleUrls: ['./filetransfertable.component.css']
})
export class FiletransfertableComponent implements OnInit, AfterViewInit {

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  @ViewChild(MatSort, {static: true}) sort: MatSort;

  projects$: Observable<Project[]>;
  employees$: Observable<Employee[]>;
  employeeUrno: number;

  displayedColumns: string[] = [ 'Jobnr', 'Project name', 'status' ];
  dataSource = new MatTableDataSource<Project>();

  constructor(
    private route: ActivatedRoute,
    private proadService: ProadService
  ) { }

  ngAfterViewInit() {
    this.projects$.subscribe(pp => this.dataSource.data = pp);
  }

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
    this.dataSource.sortingDataAccessor = (item, property) => {
      switch (property) {
        case 'Project name': return item.project_name;
        case 'Jobnr': return item.projectno;
        default: return item[property];
      }
    };
    this.projects$ = this.route.paramMap.pipe(
      switchMap((params: ParamMap) => {
        this.employeeUrno = +params.get('urno');
        return this.proadService.getProjectsForEmployee( +params.get('urno'));
      })
    );
    this.employees$ = this.proadService.getEmployees();
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

}
