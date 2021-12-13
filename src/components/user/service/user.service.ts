import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';
import { errorCode, errorService } from '../../error';

async function getUserById(id: number) {
  const userQueryRepository = getCustomRepository(UserQueryRepository);
  const user = await userQueryRepository.findOneById(id);
  if (!user) {
    throw errorService.createApolloError({
      message: 'Not Found User',
      errorCode: errorCode.BAD_REQUEST,
      params: { id },
    });
  }
  return user;
}

export default {
  getUserById,
};
