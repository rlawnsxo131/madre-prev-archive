import { FindOneOptions } from 'typeorm';
import { Data } from '.';
import { FindOneID } from '../../@types';

export default class DataService {
  static async findOne(id?: FindOneID, options?: FindOneOptions) {
    const data = await Data.findOne(id, options);
    return data;
  }
}
