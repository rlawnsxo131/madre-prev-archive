import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

function getUser(id: string) {
  return getCustomRepository(UserQueryRepository).findOneById(id);
}

export default {
  getUser,
};
