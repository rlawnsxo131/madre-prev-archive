import { Schema, string, ValidationError } from 'joi';
import ApolloCustomError from './ApolloCustomError';

export namespace ApolloValidator {
  function helper(
    error: ValidationError | undefined,
    extensions: Record<string, any>,
  ) {
    if (!error) return;
    throw new ApolloCustomError({
      message: error.message,
      code: 'BAD_REQUEST',
      extensions,
    });
  }

  export function validateId(id: string = '') {
    const schema = string().guid().required();
    const { error } = schema.validate(id);
    helper(error, { id });
  }

  export function validateObject(
    schema: Schema,
    params: Record<string, any> = {},
  ) {
    const { error } = schema.validate(params);
    if (!error) return;
    helper(error, params);
  }
}
