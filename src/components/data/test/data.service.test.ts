import 'jest';
import { Connection } from 'typeorm';
import { DataService } from '..';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';

describe('DataService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('findById user_id to 0', async () => {
    const data_id = 0;
    const data = await DataService.findById(data_id);
    expect([null, undefined]).toContain(data);
  });
});
