import { Component, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { FilterbarComponent, Filterdata } from '../filterbar/filterbar.component';
import { Project, Basecampproject } from 'src/app/model/project';
import { BasecampService } from 'src/app/services/basecamp.service';
import { Observable } from 'rxjs';


@Component({
  selector: 'app-bascamptable',
  templateUrl: './bascamptable.component.html',
  styleUrls: ['./bascamptable.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class BascamptableComponent implements OnInit, AfterViewInit {

  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;
  @ViewChild(MatSort, { static: true }) sort: MatSort;
  @ViewChild(FilterbarComponent, { static: true }) filterbar: FilterbarComponent;

  constructor(private basecampservice: BasecampService) { }

  displayedColumns: string[] = ['status', 'name', 'purpose'];
  dataSource = new MatTableDataSource<Basecampproject>();
  expandedProject: Basecampproject | null;
  basecampprojects$: Observable<Basecampproject[]>;


  ngAfterViewInit() {
   this.basecampprojects$.subscribe(pp => this.dataSource.data = pp);
  }


  ngOnInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;

    this.dataSource.filterPredicate = (project, filterstring) => {
      const splittedFilterstring = filterstring.split('#date:');
      const stringToFilter = (project.description.trim()
        + project.name.trim() + project.status.trim() + project.purpose.trim()).toLowerCase();
      if (stringToFilter.indexOf(splittedFilterstring[0]) !== -1) {
        const datestrings = splittedFilterstring[1].split(' ');
        const start = new Date(datestrings[0]);
        const end = new Date(datestrings[1]);
        const current = new Date(project.createdAt);
        if (start.getTime() <= end.getTime()) {
          return current.getTime() >= start.getTime() && current.getTime() <= end.getTime();
        }
        return false;
      }
      return false;
    };

    this.basecampprojects$ = this.basecampservice.projects();
  }

  rowClicked(project: Basecampproject) {
    this.expandedProject = this.expandedProject === project ? null : project;
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
