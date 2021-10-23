import 'jest';
import { Connection } from 'typeorm';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';

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
});
