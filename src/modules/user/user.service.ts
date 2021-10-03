import { FindOneOptions } from 'typeorm';
import { User } from '.';
import { FindOneID } from '../../@types';

export default class UserService {
  static async findOne(id?: FindOneID, options?: FindOneOptions) {
    const user = await User.findOne(id);
    return user;
  }
}
