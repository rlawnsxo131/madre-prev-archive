import { format } from 'd3';

export namespace D3Util {
  export namespace Format {
    export function formatNumberWithComma() {
      return format(',');
    }
  }

  export namespace Validation {
    export function isExistMapValidation(
      map: Map<any, any>,
      message: string = 'unknown error',
    ) {
      if (map.size) return;
      throw new Error(message);
    }
  }
}
