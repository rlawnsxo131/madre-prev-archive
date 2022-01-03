import { getCustomRepository } from 'typeorm';
import { DataQueryRepository, DataRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

export default {
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
