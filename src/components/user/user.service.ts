import { User } from '.';

export default class UserService {
  static findById(id: number) {
    return User.findOne({ id });
  }

  static findByEmail(email: string) {
    return User.findOne({ email });
  }
}
