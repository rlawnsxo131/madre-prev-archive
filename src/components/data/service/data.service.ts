import { getCustomRepository } from 'typeorm';
import { DataRepository, DataQueryRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

export function findOne(id: string) {
  return getCustomRepository(DataQueryRepository, 'default').findOneById(id);
}

export function create(params: CreateDataParams) {
  return getCustomRepository(DataRepository).createOne(params);
}
