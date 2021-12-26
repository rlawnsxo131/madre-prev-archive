import { getCustomRepository } from 'typeorm';
import { DataRepository, DataQueryRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

function getData(id: string) {
  return getCustomRepository(DataQueryRepository, 'default').findOneById(id);
}

function createData(params: CreateDataParams) {
  return getCustomRepository(DataRepository).createOne(params);
}

export default {
  getData,
  createData,
};
