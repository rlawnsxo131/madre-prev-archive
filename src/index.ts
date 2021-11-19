import { Database } from './datastore';
import { initializeEnvironment } from './lib';
import Application from './application';
import { Connection } from 'typeorm';

initializeEnvironment();
const database = new Database();
const application = new Application();

database.getConnection().then(async (_: Connection) => {
  await application.setup();
  await application.start();
});
