import 'jest';
import { Connection } from 'typeorm';
import { dataError, dataService } from '..';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';

describe('dataService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getDataById id to 0', async () => {
    const id = 0;
    try {
      await dataService.getDataById(id);
    } catch (e) {
      expect(e).toStrictEqual(dataError.notFoundData);
    }
  });
});
