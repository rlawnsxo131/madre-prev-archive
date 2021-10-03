import { FindOneOptions } from 'typeorm';
import { Data } from '.';
import { FindOneID } from '../../@types';

export default class DataService {
  static findOne(id: FindOneID, options?: FindOneOptions) {
    return Data.findOne(id, options);
  }
}
