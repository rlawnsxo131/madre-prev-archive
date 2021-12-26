import { Environment } from './lib';
import { Database } from './datastore';
import Application from './application';
import { Connection } from 'typeorm';

Environment.initialize();
const database = new Database();
const application = new Application();

database.getConnection().then(async (_: Connection) => {
  await application.setup();
  await application.start();
});
