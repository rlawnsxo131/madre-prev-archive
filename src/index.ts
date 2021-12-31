import 'reflect-metadata';
import { SetupProvider } from './lib/SetupProvider';
import { Database } from './datastore';
import Application from './application';
import { Connection } from 'typeorm';

SetupProvider.initialize();
const database = new Database();
const application = new Application();

database.getConnection().then(async (_: Connection) => {
  await application.setup();
  await application.start();
});
