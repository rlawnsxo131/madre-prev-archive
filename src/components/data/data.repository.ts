import { EntityRepository, Repository } from 'typeorm';
import { Data } from '.';

@EntityRepository(Data)
class DataRepository extends Repository<Data> {}

export default DataRepository;
