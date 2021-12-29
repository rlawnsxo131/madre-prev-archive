import { getCustomRepository } from 'typeorm';
import DataRepository from '../repository/data.repository';
import DataQueryRepository from '../repository/data.query.repository';
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
