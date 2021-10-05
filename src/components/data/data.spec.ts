import { expect } from 'chai';
import { Connection } from 'typeorm';
import { DataService } from '.';
import { Database } from '../../datastore';
import initializeEnvironment from '../../lib/initializeEnvironment';

describe('Data component tests', () => {
  initializeEnvironment();
  const database = new Database();
  let connection: Connection | null = null;

  before(async () => {
    connection = await database.getConnection();
  });

  after(async () => {
    await connection?.close();
  });

  describe('DataService', () => {
    describe('findById', () => {
      it('data_id: 0', async () => {
        const data_id = 0;
        const user = await DataService.findById(data_id);
        expect(user).to.be.oneOf([null, undefined]);
      });
    });
  });
});
