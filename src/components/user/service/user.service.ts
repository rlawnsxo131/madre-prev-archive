import { ApolloError } from 'apollo-server-core';
import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';
import { errorCode } from '../../../constants';

async function getUserById(id: number) {
  const userQueryRepository = getCustomRepository(UserQueryRepository);
  const user = await userQueryRepository.findOneById(id);
  if (!user) {
    throw new ApolloError('Not Found User', errorCode.BAD_REQUEST, { id });
  }
  return user;
}

export default {
  getUserById,
};
