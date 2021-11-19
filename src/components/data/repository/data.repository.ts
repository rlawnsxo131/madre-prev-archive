import { EntityRepository, Repository } from 'typeorm';
import { Data } from '..';

@EntityRepository(Data)
export default class DataRepository extends Repository<Data> {}
