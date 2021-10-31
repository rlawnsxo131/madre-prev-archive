import { getCustomRepository } from 'typeorm';
import { userError, UserRepository } from '.';

async function getUserById(id: number) {
  const userRepo = getCustomRepository(UserRepository);
  const user = await userRepo.findOne({ id });
  if (!user) {
    throw userError.notFoundUser;
  }
  return user;
}

export default {
  getUserById,
};
