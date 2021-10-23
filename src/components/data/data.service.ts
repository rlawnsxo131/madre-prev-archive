import { getCustomRepository } from 'typeorm';
import { DataRepository } from '.';

function findById(id: number) {
  const dataRepo = getCustomRepository(DataRepository);
  return dataRepo.findOne(id);
}

export default {
  findById,
};
