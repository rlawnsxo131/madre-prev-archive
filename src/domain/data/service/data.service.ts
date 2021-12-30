import { getCustomRepository } from 'typeorm';
import { DataQueryRepository, DataRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

function findOne(id: string) {
  return getCustomRepository(DataQueryRepository, 'default').findOneById(id);
}

function create(params: CreateDataParams) {
  return getCustomRepository(DataRepository).createOne(params);
}

const DataService = {
  findOne,
  create,
};

export default DataService;
