import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';

function findById(id: number) {
  const userRepo = getCustomRepository(UserRepository);
  return userRepo.findOne({ id });
}

export default {
  findById,
};
