import 'jest';
import { ApolloError } from 'apollo-server-core';
import { apolloErrorService } from '..';

describe('apolloErrorService Test', () => {
  test('throwErrorValidation: type NOT_FOUND and id to 1', async () => {
    const id = 1;
    const code = 'NOT_FOUND';
    try {
      apolloErrorService.throwErrorValidation({
        resolver: () => true,
        message: 'Not Found Data',
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
