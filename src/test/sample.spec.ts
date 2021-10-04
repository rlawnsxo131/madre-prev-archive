import { expect } from 'chai';
import { Database } from '../db';
import initializeEnvironment from '../lib/initializeEnvironment';

initializeEnvironment();
const database = new Database();

describe('sample test', () => {
  it('hello world', () => {
    const world = 'world';
    expect(world).equal('world');
  });

  it('database connection test', async () => {
    const connection = await database.getConnection();
    expect(connection).not.null;
  });
});
