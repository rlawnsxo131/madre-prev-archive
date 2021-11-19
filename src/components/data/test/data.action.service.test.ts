import 'jest';
import { ApolloError } from 'apollo-server-core';
import { Connection } from 'typeorm';
import { dataActionService } from '..';
import { Database } from '../../../datastore';
import { initializeEnvironment } from '../../../lib';
import { errorCode } from '../../error';

describe('dataActionService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getData id to 0', async () => {
    const id = 0;
    try {
      await dataActionService.getDataAction(id);
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(errorCode.NOT_FOUND);
      expect(error.extensions.id).toBe(id);
    }
  });
});
