import { ApolloError } from 'apollo-server-core';
import 'jest';
import { Connection } from 'typeorm';
import { userService } from '..';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';
import { errorCode } from '../../error';

describe('userService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getUserById id to 0', async () => {
    const id = 0;
    try {
      await userService.getUserById(id);
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(errorCode.NOT_FOUND);
      expect(error.extensions.id).toBe(id);
    }
  });
});
