import { getCustomRepository } from 'typeorm';
import UserQueryRepository from '../repository/user.query.repository';

function findOne(id: string) {
  return getCustomRepository(UserQueryRepository, 'default').findOneById(id);
}

const UserService = {
  findOne,
};

export default UserService;
