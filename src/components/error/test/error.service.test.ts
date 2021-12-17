import 'jest';
import { errorService } from '..';
import { ApolloError } from 'apollo-server-core';

describe('errorService Test', () => {
  test('throwApolloError: id to 1', async () => {
    const id = 1;
    try {
      errorService.throwApolloError({
        resolver: () => true,
        message: 'Not Found Data',
        code: 'NOT_FOUND',
        params: { id },
      });
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(errorService.ERROR_CODE.NOT_FOUND);
      expect(error.extensions.id).toBe(id);
    }
  });
});
