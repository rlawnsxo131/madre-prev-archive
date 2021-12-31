import { Service } from 'typedi';
import { EntityRepository, Repository } from 'typeorm';
import { Data } from '..';
import { CreateDataParams } from '../interface/data.interface';

@Service()
@EntityRepository(Data)
export default class DataRepository extends Repository<Data> {
  createOne(params: CreateDataParams) {
    const data = this.create(params);
    return this.save(data);
  }
}
