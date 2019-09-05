import { Component, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { ProadService, StatusCode } from 'src/app/services/proad.service';
import { Project } from 'src/app/model/project';
import { Observable } from 'rxjs';
import { Employee } from 'src/app/model/employee';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { FilterbarComponent, Filterdata } from '../filterbar/filterbar.component';
import { animate, state, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'app-filetransfertable',
  templateUrl: './filetransfertable.component.html',
  styleUrls: ['./filetransfertable.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class FiletransfertableComponent implements OnInit, AfterViewInit {

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  @ViewChild(MatSort, {static: true}) sort: MatSort;
  @ViewChild(FilterbarComponent, {static: true}) filterbar: FilterbarComponent;

  projects$: Observable<Project[]>;
  employees$: Observable<Employee[]>;
  employeeUrno: number;

  displayedColumns: string[] = [ 'Jobnr', 'Projektname', 'Auftragsdatum', 'Status' ];
  dataSource = new MatTableDataSource<Project>();
  expandedProject: Project | null;

  constructor(
    private proadService: ProadService,
  ) { }

  ngAfterViewInit() {
    this.projects$.subscribe(pp => this.dataSource.data = pp);
  }

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
    this.dataSource.sortingDataAccessor = (item, property) => {
      switch (property) {
        case 'Projektname': return item.project_name;
        case 'Jobnr': return item.projectno;
        case 'Auftragsdatum': return item.order_date;
        default: return item[property];
      }
    };
    this.dataSource.filterPredicate = (project, filterstring) => {
      const splittedFilterstring = filterstring.split('#date:');
      const stringToFilter = (project.description.trim()
      + project.project_name.trim() + project.status.trim() + project.projectno.trim()).toLowerCase();
      if (stringToFilter.indexOf(splittedFilterstring[0]) !== -1) {
        const datestrings = splittedFilterstring[1].split(' ');
        const start = new Date(datestrings[0]);
        const end = new Date(datestrings[1]);
        const current = new Date(project.order_date);
        if (start.getTime() <= end.getTime()) {
          return current.getTime() >= start.getTime() && current.getTime() <= end.getTime();
        }
        return false;
      }
      return false;
    };
    const eDate = new Date();
    const sDate = new Date();
    sDate.setMonth(0);
    this.filterbar.startDate = sDate;
    this.filterbar.endDate = eDate;
    this.projects$ = this.proadService.getFilteredProjects(StatusCode.none, sDate, eDate);
    this.employees$ = this.proadService.getEmployees();
  }

  filterChanged(filterdata: Filterdata) {
    const filterstring = filterdata.text.trim().toLowerCase()
    + '#date:'
    + filterdata.startDate.toISOString()
    + ' '
    + filterdata.endDate.toISOString();
    console.log(filterstring);
    this.dataSource.filter = filterstring;
  }
}
