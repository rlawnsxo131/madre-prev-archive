import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

function getOneById(id: number) {
  return getCustomRepository(UserQueryRepository).findOneById(id);
}

export default {
  getOneById,
};
