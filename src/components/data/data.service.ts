import { getCustomRepository } from 'typeorm';
import { DataRepository } from '.';

const dataService = {
  findById(id: number) {
    const dataRepo = getCustomRepository(DataRepository);
    return dataRepo.findOne(id);
  },
};

export default dataService;
