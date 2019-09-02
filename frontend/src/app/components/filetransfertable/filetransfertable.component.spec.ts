import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FiletransfertableComponent } from './filetransfertable.component';

describe('FiletransfertableComponent', () => {
  let component: FiletransfertableComponent;
  let fixture: ComponentFixture<FiletransfertableComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FiletransfertableComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FiletransfertableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
