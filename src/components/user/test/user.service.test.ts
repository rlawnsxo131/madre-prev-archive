import 'jest';
import { Environment } from '../../../lib';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { UserService } from '..';

describe('UserService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    Environment.initialize();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getUser: id to undefined', async () => {
    const id = 'undefined';
    const user = await UserService.getUser(id);
    expect(user).toBe(undefined);
  });
});
