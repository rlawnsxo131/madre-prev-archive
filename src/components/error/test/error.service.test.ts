import 'jest';
import { errorService } from '..';
import { ApolloError } from 'apollo-server-core';

describe('errorService Test', () => {
  test('throwApolloError: type NOT_FOUND and id to 1', async () => {
    const id = 1;
    const code = 'NOT_FOUND';
    try {
      errorService.throwApolloError({
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
