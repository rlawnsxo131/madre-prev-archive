import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

function getUserById(id: number) {
  return getCustomRepository(UserQueryRepository).findOneById(id);
}

export default {
  getUserById,
};
