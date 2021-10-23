import { getCustomRepository } from 'typeorm';
import { DataRepository } from '.';

export default {
  findById(id: number) {
    const dataRepo = getCustomRepository(DataRepository);
    return dataRepo.findOne(id);
  },
};
