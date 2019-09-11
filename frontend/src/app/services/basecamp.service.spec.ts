import { TestBed } from '@angular/core/testing';

import { BasecampService } from './basecamp.service';

describe('BasecampService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: BasecampService = TestBed.get(BasecampService);
    expect(service).toBeTruthy();
  });
});
