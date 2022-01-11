import { EntityRepository, Repository } from 'typeorm';
import { Data } from '..';

@EntityRepository(Data)
export default class DataQueryRepository extends Repository<Data> {
  findAll() {
    return this.createQueryBuilder('data').getMany();
  }

  findOneById(id: string) {
    return this.createQueryBuilder('data')
      .select('data')
      .where('data.id = :id', { id })
      .getOne();
  }
}
