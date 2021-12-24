import { getCustomRepository } from 'typeorm';
import { DataRepository, DataQueryRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

export namespace DataService {
  export function getData(id: string) {
    return getCustomRepository(DataQueryRepository, 'default').findOneById(id);
  }

  export function createData(params: CreateDataParams) {
    return getCustomRepository(DataRepository).createOne(params);
  }
}
