import FastifyCustomError from '../FastifyCustomError';
import { validateId, validateObject } from './validator';
import { Validator, ValidatorHelper } from './validator.type';

const fastifyHelper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new FastifyCustomError({
    message: `params: ${JSON.stringify(params)}\n${error.message}`,
    name: 'BadRequestError',
    statusCode: 400,
  });
};

const fastifyValidator: Validator = {
  validateId: (id) => validateId(id)(fastifyHelper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(fastifyHelper),
};

export default fastifyValidator;
