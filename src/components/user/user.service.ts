import { getCustomRepository } from 'typeorm';
import { UserRepository } from '.';
import { errorCode, errorService } from '../error';

async function getUserById(id: number) {
  const userRepo = getCustomRepository(UserRepository);
  const user = await userRepo.findOne({ id });
  if (!user) {
    throw errorService.createApolloError({
      message: 'Not Found User',
      errorCode: errorCode.NOT_FOUND,
      params: { id },
    });
  }
  return user;
}

export default {
  getUserById,
};
