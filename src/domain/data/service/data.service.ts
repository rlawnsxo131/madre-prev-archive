import { getCustomRepository } from 'typeorm';
import { DataQueryRepository, DataRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

const DataService = {
  findAll() {
    return getCustomRepository(DataQueryRepository).findAll();
  },
  findOne(id: string) {
    return getCustomRepository(DataQueryRepository).findOneById(id);
  },
  create(params: CreateDataParams) {
    return getCustomRepository(DataRepository).createOne(params);
  },
};

export default DataService;
