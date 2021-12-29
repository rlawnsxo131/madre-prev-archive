import { config } from 'dotenv';
import { resolve } from 'path';
import { environmentFilename } from '../constants';

export namespace SetupProvider {
  export function initialize() {
    return config({
      path: resolve(process.cwd(), environmentFilename),
    });
  }
}
