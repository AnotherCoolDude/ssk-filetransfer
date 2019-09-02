import { TestBed } from '@angular/core/testing';

import { ProadService } from './proad.service';

describe('ProadService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ProadService = TestBed.get(ProadService);
    expect(service).toBeTruthy();
  });
});
