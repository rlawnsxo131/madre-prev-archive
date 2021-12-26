import { config } from 'dotenv';
import { resolve } from 'path';
import { environmentFilename } from '../constants';

function initialize() {
  config({
    path: resolve(process.cwd(), environmentFilename),
  });
}

export default {
  initialize,
};
