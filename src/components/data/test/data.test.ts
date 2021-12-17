import 'jest';
import { initializeEnvironment } from '../../../lib';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { dataService } from '..';
import { errorService } from '../../error';
import { ApolloError } from 'apollo-server-core';

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

  test('getData: id to 0', async () => {
    const id = 0;
    const data = await dataService.getData(id);
    expect(data).toBe(undefined);
  });

  test('getData: id to 0 and throwApolloError', async () => {
    const id = 0;
    try {
      const data = await dataService.getData(id);
      errorService.throwApolloError({
        resolver: () => !data,
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
