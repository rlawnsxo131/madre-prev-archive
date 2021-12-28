import { Schema } from 'joi';

export namespace JoiValidator {
  interface ValidateParams {
    schema: Schema;
    params: any;
  }

  export function validate({ schema, params }: ValidateParams) {
    const { error } = schema.validate(params);
    return error;
  }
}
