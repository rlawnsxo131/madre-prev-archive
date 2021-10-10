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
    const userRepo = getCustomRepository(UserRepository);
    return userRepo.findOne({ id });
  }
}

export default UserService.getInstance();
