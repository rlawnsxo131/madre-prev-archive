import 'jest';
import { SetupProvider } from '../../../lib/SetupProvider';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { UserService } from '..';
import Container from 'typedi';

describe('UserService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    SetupProvider.initialize();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('findOne: id to undefined', async () => {
    const id = 'undefined';
    const user = await Container.get(UserService).findOne(id);
    expect(user).toBe(undefined);
  });
});
