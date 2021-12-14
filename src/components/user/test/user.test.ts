import 'jest';
import { ApolloError } from 'apollo-server-core';
import { Connection } from 'typeorm';
import { userService } from '..';
import { Database } from '../../../datastore';
import { initializeEnvironment } from '../../../lib';
import { ERROR_CODE } from '../../../constants';

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

  test('getUser: id to 0 and throw', async () => {
    const id = 0;
    try {
      await userService.getUser(id);
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(ERROR_CODE.BAD_REQUEST);
      expect(error.extensions.id).toBe(id);
    }
  });
});
