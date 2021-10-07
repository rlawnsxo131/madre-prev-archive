import { EntityRepository, Repository } from 'typeorm';
import { User } from '.';

@EntityRepository(User)
class UserRepository extends Repository<User> {}

export default UserRepository;
