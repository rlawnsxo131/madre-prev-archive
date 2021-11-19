import 'jest';
import { Connection } from 'typeorm';
import { Database } from '../../../datastore';
import { initializeEnvironment } from '../../../lib';

describe('authService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
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
