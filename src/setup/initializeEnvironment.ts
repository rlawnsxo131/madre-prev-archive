import { config } from 'dotenv';
import { resolve } from 'path';
import { environmentFilename } from '../constants';
import logger from '../lib/logger/logger';

export default function initializeEnvironment() {
  logger.info('run initalizeEnvironment');
  return config({
    path: resolve(process.cwd(), environmentFilename),
  });
}
