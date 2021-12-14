import { config } from 'dotenv';
import { resolve } from 'path';
import { ENVIRONMENT_FILENAME } from '../constants';

export default function initializeEnvironment() {
  config({
    path: resolve(process.cwd(), ENVIRONMENT_FILENAME),
  });
}
