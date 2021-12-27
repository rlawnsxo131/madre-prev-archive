import { config } from 'dotenv';
import { resolve } from 'path';
import { environmentFilename } from '../constants';

// class SingletonSetupProvider {
//   private static instance: SingletonSetupProvider;
//   private constructor() {}

//   static getInstance() {
//     if (!this.instance) {
//       this.instance = new SingletonSetupProvider();
//     }
//     return this.instance;
//   }

//   initialize() {
//     return config({
//       path: resolve(process.cwd(), environmentFilename),
//     });
//   }
// }

// const SetupProvider = SingletonSetupProvider.getInstance();

// export default SetupProvider;

export namespace SetupProvider {
  export function initialize() {
    return config({
      path: resolve(process.cwd(), environmentFilename),
    });
  }
}
