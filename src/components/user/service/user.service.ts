import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

function getUser(id: number) {
  return getCustomRepository(UserQueryRepository).findOneById(id);
}

export default {
  getUser,
};
