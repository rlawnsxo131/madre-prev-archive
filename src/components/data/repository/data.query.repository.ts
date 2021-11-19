import { EntityRepository, Repository } from 'typeorm';
import { Data } from '..';

@EntityRepository(Data)
export default class DataQueryRepository extends Repository<Data> {
  findOneById(id: number) {
    return this.createQueryBuilder('data')
      .select('data')
      .where('data.id = :id', { id })
      .getOne();
  }
}
