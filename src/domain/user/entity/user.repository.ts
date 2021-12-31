import { Service } from 'typedi';
import { EntityRepository, Repository } from 'typeorm';
import { User } from '..';

@Service()
@EntityRepository(User)
export default class UserRepository extends Repository<User> {}
