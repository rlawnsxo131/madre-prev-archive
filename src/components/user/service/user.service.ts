import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

function getUser(id: string) {
  return getCustomRepository(UserQueryRepository, 'default').findOneById(id);
}

export default {
  getUser,
};
