import { EntityRepository, Repository } from 'typeorm';
import { Data } from '..';

/**
 * CREATE, UPDATE, DELETE
 */
@EntityRepository(Data)
export default class DataRepository extends Repository<Data> {}
