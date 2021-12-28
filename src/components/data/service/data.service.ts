import { getCustomRepository } from 'typeorm';
import { DataRepository, DataQueryRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

function findOne(id: string) {
  return getCustomRepository(DataQueryRepository, 'default').findOneById(id);
}

function create(params: CreateDataParams) {
  return getCustomRepository(DataRepository).createOne(params);
}

const userService = {
  findOne,
  create,
};

export default userService;
