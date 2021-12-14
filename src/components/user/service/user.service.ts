import { ApolloError } from 'apollo-server-core';
import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';
import { ERROR_CODE } from '../../../constants';

async function getUserById(id: number) {
  const userQueryRepository = getCustomRepository(UserQueryRepository);
  const user = await userQueryRepository.findOneById(id);
  if (!user) {
    throw new ApolloError('Not Found User', ERROR_CODE.BAD_REQUEST, { id });
  }
  return user;
}

export default {
  getUserById,
};
