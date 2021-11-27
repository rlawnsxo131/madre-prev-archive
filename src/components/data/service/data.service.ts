import { getCustomRepository } from 'typeorm';
import DataQueryRepository from '../repository/data.query.repository';

function getOneById(id: number) {
  return getCustomRepository(DataQueryRepository).findOneById(id);
}

export default {
  getOneById,
};
