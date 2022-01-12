import { validateId, validateObject } from './validator';
import ApolloCustomError from '../ApolloCustomError';
import { Validator, ValidatorHelper } from './validation';

const helper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new ApolloCustomError({
    message: error.message,
    code: 'BAD_REQUEST',
    extensions: params,
  });
};

const apolloValidator: Validator = {
  validateId: (id) => validateId(id)(helper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(helper),
};

export default apolloValidator;
