import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

export namespace UserService {
  export function getUser(id: string) {
    return getCustomRepository(UserQueryRepository, 'default').findOneById(id);
  }
}
