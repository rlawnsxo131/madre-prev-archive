import 'jest';
import { SetupProvider } from '../../../lib/SetupProvider';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { userService } from '..';

describe('userService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    SetupProvider.initialize();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getUser: id to undefined', async () => {
    const id = 'undefined';
    const user = await userService.getUser(id);
    expect(user).toBe(undefined);
  });
});
