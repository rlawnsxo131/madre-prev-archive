import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

export function getUser(id: string) {
  return getCustomRepository(UserQueryRepository, 'default').findOneById(id);
}
