import { config } from 'dotenv';
import { resolve } from 'path';
import constant from '../constant';

export default function initializeEnvironment() {
  config({
    path: resolve(process.cwd(), constant.environmentFilename),
  });
}
