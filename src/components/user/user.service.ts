import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';

const userService = {
  findById(id: number) {
    const userRepo = getCustomRepository(UserRepository);
    return userRepo.findOne({ id });
  },
};

export default userService;
