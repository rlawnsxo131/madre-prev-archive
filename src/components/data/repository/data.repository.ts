import { DeepPartial, EntityRepository, Repository } from 'typeorm';
import { Data } from '..';

@EntityRepository(Data)
export default class DataRepository extends Repository<Data> {
  createOne(values: DeepPartial<Data>) {
    const data = this.create(values);
    return this.save(data);
  }
}
