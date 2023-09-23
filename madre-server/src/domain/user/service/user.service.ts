import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

export default {
  findOne(id: string) {
    return getCustomRepository(UserQueryRepository).findOneById(id);
  },
};
