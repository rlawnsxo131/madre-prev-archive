import 'jest';
import { Connection } from 'typeorm';
import { authService } from '..';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../setup/initializeEnvironment';

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
    authService.authFunction();
  });
});
