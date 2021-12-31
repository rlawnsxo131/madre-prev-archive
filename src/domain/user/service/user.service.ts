import { Service } from 'typedi';
import { InjectRepository } from 'typeorm-typedi-extensions';
import { UserQueryRepository } from '..';

@Service()
export default class UserService {
  constructor(
    @InjectRepository(UserQueryRepository, 'default')
    private readonly userQueryRepository: UserQueryRepository,
  ) {}

  findOne(id: string) {
    return this.userQueryRepository.findOneById(id);
  }
}
