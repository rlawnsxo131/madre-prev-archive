import { config } from 'dotenv';
import { resolve } from 'path';
import { environmentFilename } from '../constants';

export namespace Environment {
  export function initialize() {
    config({
      path: resolve(process.cwd(), environmentFilename),
    });
  }
}
