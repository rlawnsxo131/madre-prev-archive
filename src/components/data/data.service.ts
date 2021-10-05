import { Data } from '.';

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
    return Data.findOne({ id });
  }
}

const dataService = DataService.getInstance();

export default dataService;
