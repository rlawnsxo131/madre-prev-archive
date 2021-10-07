import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';

class UserService {
  private static instance: UserService;
  private userRepository = getCustomRepository(UserRepository);

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new UserService();
    }
    return this.instance;
  }

  findById(id: number) {
    return this.userRepository.findOne({ id });
  }

  findByEmail(email: string) {
    return this.userRepository.findOne({ email });
  }
}

export default UserService.getInstance();
