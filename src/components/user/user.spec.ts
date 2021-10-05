import { expect } from 'chai';
import { Connection } from 'typeorm';
import { UserService } from '.';
import { Database } from '../../datastore';
import initializeEnvironment from '../../lib/initializeEnvironment';

describe('User components tests', () => {
  initializeEnvironment();
  const database = new Database();
  let connection: Connection | null = null;

  before(async () => {
    connection = await database.getConnection();
  });

  after(async () => {
    await connection?.close();
  });

  describe('UserService', () => {
    describe('findById', () => {
      it('user_id: 0', async () => {
        const user_id = 0;
        const user = await UserService.findById(user_id);
        expect(user).to.be.oneOf([null, undefined]);
      });
    });
  });
});
