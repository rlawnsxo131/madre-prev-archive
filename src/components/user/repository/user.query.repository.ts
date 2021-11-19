import { EntityRepository, Repository } from 'typeorm';
import { User } from '..';

@EntityRepository(User)
export default class UserQueryRepository extends Repository<User> {
  findOneById(id: number) {
    return this.createQueryBuilder('user')
      .select('user')
      .where('user.id = :id', { id })
      .getOne();
  }
}
