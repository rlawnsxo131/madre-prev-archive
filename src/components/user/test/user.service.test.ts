import 'jest';
import { initializeEnvironment } from '../../../lib';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { userService } from '..';

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

  test('getUser: id to 0', async () => {
    const id = 0;
    const user = await userService.getUser(id);
    expect(user).toBe(undefined);
  });
});
