import { Database } from './datastore';
import initializeEnvironment from './lib/initializeEnvironment';
import Application from './application';

initializeEnvironment();
const database = new Database();
const application = new Application();

database.getConnection().then(async () => {
  await application.setup();
  await application.start();
});
