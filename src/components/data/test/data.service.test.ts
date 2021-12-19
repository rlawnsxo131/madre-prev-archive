import 'jest';
import { initializeEnvironment } from '../../../lib';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { dataService } from '..';
import { CreateDataParams } from '../interface/data.interface';

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

  test('getData: id to undefined', async () => {
    const id = 'undefined';
    const data = await dataService.getData(id);
    expect(data).toBe(undefined);
  });

  test('createData', async () => {
    const userObj: CreateDataParams = {
      user_id: 'user_id',
      file_url: 'asdf',
      title: 'title',
      description: undefined,
      is_public: false,
    };
    const data = await dataService.createData(userObj);
    expect(data.user_id).toBe(userObj.user_id);
    expect(data.file_url).toBe(userObj.file_url);
    expect(data.title).toBe(userObj.title);
    expect(data.description).toBe(null);
    expect(data.is_public).toBe(userObj.is_public);
  });
});
