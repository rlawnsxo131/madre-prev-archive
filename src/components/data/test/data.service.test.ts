import 'jest';
import { initializeEnvironment } from '../../../lib';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { dataService } from '..';

describe('dataService Test', () => {
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
});
