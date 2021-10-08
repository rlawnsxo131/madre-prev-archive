import { Service } from 'typedi';
import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';

@Service()
class UserService {
  private static instance: UserService;

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new UserService();
    }
    return this.instance;
  }

  async findById(id: number) {
    return getCustomRepository(UserRepository).findOne({ id });
  }
}

export default UserService.getInstance();
