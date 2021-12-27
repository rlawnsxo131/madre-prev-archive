import { getCustomRepository } from 'typeorm';
import { UserQueryRepository } from '..';

export function findOne(id: string) {
  return getCustomRepository(UserQueryRepository, 'default').findOneById(id);
}
