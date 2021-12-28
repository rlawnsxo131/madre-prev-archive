import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

function findOne(id: string) {
  return getCustomRepository(UserQueryRepository, 'default').findOneById(id);
}

const userService = {
  findOne,
};

export default userService;
