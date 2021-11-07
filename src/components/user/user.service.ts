import { getCustomRepository } from 'typeorm';
import { userError, UserRepository } from '.';
import { errorCode, errorService } from '../error';

async function getUserById(id: number) {
  const userRepo = getCustomRepository(UserRepository);
  const user = await userRepo.findOne({ id });
  if (!user) {
    throw errorService.createError({
      message: userError.errorMessage.notFoundUser,
      errorCode: errorCode.NOT_FOUND,
      params: { id },
    });
  }
  return user;
}

export default {
  getUserById,
};
