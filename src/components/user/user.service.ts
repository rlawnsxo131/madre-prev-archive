import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';

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
    return getCustomRepository(UserRepository).findOne({ id });
  }

  findByEmail(email: string) {
    return getCustomRepository(UserRepository).findOne({ email });
  }
}

export default UserService.getInstance();
