import { EntityRepository, Repository } from 'typeorm';
import { User } from '..';

@EntityRepository(User)
export class UserRepository extends Repository<User> {}
