import { getCustomRepository, getRepository } from 'typeorm';
import { DataRepository } from '.';

class DataService {
  private static instance: DataService;

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new DataService();
    }
    return this.instance;
  }

  findById(id: number) {
    const dataRepo = getRepository(DataRepository);
    return dataRepo.findOne(id);
  }
}

export default DataService.getInstance();
