import { validateId, validateObject } from './validator';
import FastifyCustomError from '../FastifyCustomError';
import { Validator, ValidatorHelper } from './validation';

const helper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new FastifyCustomError({
    message: `params: ${JSON.stringify(params)}\n${error.message}`,
    name: 'BadRequestError',
    statusCode: 400,
  });
};

const fastifyValidator: Validator = {
  validateId: (id) => validateId(id)(helper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(helper),
};

export default fastifyValidator;
