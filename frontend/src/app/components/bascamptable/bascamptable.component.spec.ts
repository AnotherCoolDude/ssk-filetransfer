import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BascamptableComponent } from './bascamptable.component';

describe('BascamptableComponent', () => {
  let component: BascamptableComponent;
  let fixture: ComponentFixture<BascamptableComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BascamptableComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BascamptableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
