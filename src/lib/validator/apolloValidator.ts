import ApolloCustomError from '../ApolloCustomError';
import { validateId, validateObject } from './validator';
import { Validator, ValidatorHelper } from './validator.type';

const apolloHelper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new ApolloCustomError({
    message: error.message,
    code: 'BAD_REQUEST',
    extensions: params,
  });
};

const apolloValidator: Validator = {
  validateId: (id) => validateId(id)(apolloHelper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(apolloHelper),
};

export default apolloValidator;
