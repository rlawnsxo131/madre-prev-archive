import { format } from 'd3';

export default class D3Util {
  static formatNumberWithComma() {
    return format(',');
  }

  static isExistMapValidation(
    map: Map<any, any>,
    message: string = 'unknown error',
  ) {
    if (map.size) return;
    throw new Error(message);
  }
}
