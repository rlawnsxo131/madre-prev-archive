import { getCustomRepository } from 'typeorm';
import { DataRepository } from '.';

class DataService {
  private static instance: DataService;
  private dataRepository = getCustomRepository(DataRepository);

  private constructor() {}

  static getInstance() {
    if (!this.instance) {
      this.instance = new DataService();
    }
    return this.instance;
  }

  findById(id: number) {
    return this.dataRepository.findOne({ id });
  }
}

export default DataService.getInstance();
