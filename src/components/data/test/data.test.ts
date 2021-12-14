import 'jest';
import { Connection } from 'typeorm';
import { Database } from '../../../datastore';
import { initializeEnvironment } from '../../../lib';
import { dataService } from '..';
import { ApolloError } from 'apollo-server-core';
import { ERROR_CODE } from '../../../constants';

describe('data Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getData: id to 0 and throw', async () => {
    const id = 0;
    try {
      const data = await dataService.getData(id);
      if (!data) {
        throw new ApolloError('Not Found Data', ERROR_CODE.NOT_FOUND, { id });
      }
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(ERROR_CODE.NOT_FOUND);
      expect(error.extensions.id).toBe(id);
    }
  });
});
