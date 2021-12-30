import { getCustomRepository } from 'typeorm';
import UserQueryRepository from '../repository/user.query.repository';

const UserService = {
  findOne(id: string) {
    return getCustomRepository(UserQueryRepository).findOneById(id);
  },
};

export default UserService;
