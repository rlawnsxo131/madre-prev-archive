import { FindOneOptions } from 'typeorm';
import { User } from '.';
import { FindOneID } from '../../@types';

export default class UserService {
  static findOne(id: FindOneID, options?: FindOneOptions) {
    return User.findOne(id);
  }
}
