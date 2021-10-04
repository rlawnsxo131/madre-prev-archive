import { Data } from '.';

export default class DataService {
  static findById(id: number) {
    return Data.findOne({ id });
  }
}
