import { EntityRepository, Repository } from 'typeorm';
import { User } from '..';

@EntityRepository(User)
export class UserQueryRepository extends Repository<User> {
  findOneById(id: string) {
    return this.createQueryBuilder('user')
      .select('user')
      .where('user.id = :id', { id })
      .getOne();
  }
}
