import { EntityRepository, Repository } from 'typeorm';
import { User } from '..';

@EntityRepository(User)
export default class UserRepository extends Repository<User> {}
