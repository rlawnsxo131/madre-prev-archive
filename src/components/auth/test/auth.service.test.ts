import 'jest';
import { environmentManager } from '../../../lib';
import { Connection } from 'typeorm';
import { Database } from '../../../datastore';

describe('authService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    environmentManager.initialize();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('auth', () => {
    console.log('hello');
  });
});
