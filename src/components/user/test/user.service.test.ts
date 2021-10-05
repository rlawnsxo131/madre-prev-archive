import 'jest';
import { Connection } from 'typeorm';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';
import { UserService } from '..';

describe('UserService Test', () => {
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
    const user_id = 0;
    const user = await UserService.findById(user_id);
    expect([null, undefined]).toContain(user);
  });
});
