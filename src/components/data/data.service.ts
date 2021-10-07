import { getCustomRepository } from 'typeorm';
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
    return getCustomRepository(DataRepository).findOne({ id });
  }
}

const dataService = DataService.getInstance();

export default dataService;
