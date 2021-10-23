import 'jest';
import { Connection } from 'typeorm';
import { userService } from '..';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';

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

  test('findById user_id to 0', async () => {
    const user_id = 0;
    const user = await userService.findById(user_id);
    expect([null, undefined]).toContain(user);
  });
});
