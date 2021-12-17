import 'jest';
import { initializeEnvironment } from '../../../lib';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { userService } from '..';
import { errorService } from '../../error';
import { ApolloError } from 'apollo-server-core';

describe('user Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getUser: id to 0', async () => {
    const id = 0;
    const user = await userService.getUser(id);
    expect(user).toBe(undefined);
  });

  test('getUser: id to 0 with throwApolloError', async () => {
    const id = 0;
    try {
      const user = await userService.getUser(id);
      errorService.throwApolloError({
        resolver: () => !user,
        message: 'Not Found User',
        code: 'BAD_REQUEST',
        params: { id },
      });
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(errorService.ERROR_CODE.BAD_REQUEST);
      expect(error.extensions.id).toBe(id);
    }
  });
});
