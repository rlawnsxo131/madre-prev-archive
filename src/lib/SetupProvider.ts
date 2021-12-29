import { config } from 'dotenv';
import { resolve } from 'path';
import { environmentFilename } from '../constants';

class SetupProvider {
  constructor() {}

  initialize() {
    return config({
      path: resolve(process.cwd(), environmentFilename),
    });
  }
}

const setupProvider = new SetupProvider();

export default setupProvider;
