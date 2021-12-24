import 'jest';
import { UserValidationService } from '..';

describe('userValidationService Test', () => {
  test('getUserParamsValidation', () => {
    const id = '';
    try {
      UserValidationService.getUserParamsValidation(id);
    } catch (e: any) {
      expect(e.extensions.code).toBe('BAD_REQUEST');
      expect(e.extensions.id).toBe(id);
    }
  });
});
