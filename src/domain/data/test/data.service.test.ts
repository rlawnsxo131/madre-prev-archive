import 'jest';
import { SetupProvider } from '../../../lib/SetupProvider';
import { Database } from '../../../datastore';
import { Connection } from 'typeorm';
import { DataService } from '..';
import { CreateDataParams } from '../interface/data.interface';

describe('DataService Test', () => {
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
    const data = await DataService.findOne(id);
    expect(data).toBe(undefined);
  });

  test('create', async () => {
    const createDataParams: CreateDataParams = {
      user_id: 'cc950501-63ff-11ec-95c1-0242ac130002',
      file_url: 'asdf',
      title: 'title',
      description: undefined,
      is_public: false,
    };
    const data = await DataService.create(createDataParams);
    expect(data.user_id).toBe(createDataParams.user_id);
    expect(data.file_url).toBe(createDataParams.file_url);
    expect(data.title).toBe(createDataParams.title);
    expect(data.description).toBe(null);
    expect(data.is_public).toBe(createDataParams.is_public);
  });
});
