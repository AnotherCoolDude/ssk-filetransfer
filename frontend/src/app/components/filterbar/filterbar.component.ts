import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-filterbar',
  templateUrl: './filterbar.component.html',
  styleUrls: ['./filterbar.component.css']
})

export class FilterbarComponent implements OnInit {

  filterText: string;
  startDate: Date;
  endDate: Date;

  @Output() filterdatachanged: EventEmitter<Filterdata> = new EventEmitter();

  constructor() { }

  ngOnInit() {  }

  filterdataChanged() {
    this.filterdatachanged.emit(new Filterdata(this.filterText, this.startDate, this.endDate));
  }

}

export class Filterdata {
  text: string;
  startDate: Date;
  endDate: Date;

  constructor(text: string, sDate: Date, eDate: Date) {
    this.text = text;
    this.startDate = sDate;
    this.endDate = eDate;
  }
}
