import 'jest';
import { ApolloError } from 'apollo-server-core';
import { errorService } from '../../error';
import { userValidationService } from '..';

describe('userValidationService Test', () => {
  test('getUserParamsValidation', async () => {
    const id = '';
    const code = 'BAD_USER_INPUT';
    try {
      const validation = userValidationService.getUserParamsValidation(id);
      errorService.throwApolloError({
        resolver: () => !validation,
        message: 'validation test',
        code,
        params: { id },
      });
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(code);
      expect(error.extensions.id).toBe(id);
    }
  });
});
