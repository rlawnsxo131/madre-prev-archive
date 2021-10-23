import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';

export default {
  findById(id: number) {
    const userRepo = getCustomRepository(UserRepository);
    return userRepo.findOne({ id });
  },
};
