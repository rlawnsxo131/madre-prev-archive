import { getCustomRepository } from 'typeorm';
import DataQueryRepository from '../repository/data.query.repository';

async function getData(id: number) {
  return getCustomRepository(DataQueryRepository).findOneById(id);
}

export default {
  getData,
};
