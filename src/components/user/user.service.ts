import { User } from '.';

class UserService {
  private static instance: UserService;

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new UserService();
    }
    return this.instance;
  }

  findById(id: number) {
    return User.findOne({ id });
  }

  findByEmail(email: string) {
    return User.findOne({ email });
  }
}

const userService = UserService.getInstance();

export default userService;
