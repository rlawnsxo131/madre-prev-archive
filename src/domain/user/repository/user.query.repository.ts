import { Service } from 'typedi';
import { EntityRepository, Repository } from 'typeorm';
import { User } from '..';

@Service()
@EntityRepository(User)
export default class UserQueryRepository extends Repository<User> {
  findOneById(id: string) {
    return this.createQueryBuilder('user')
      .select('user')
      .where('user.id = :id', { id })
      .getOne();
  }
}
