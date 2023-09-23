import ApolloCustomError from '../ApolloCustomError';
import FastifyCustomError from '../FastifyCustomError';
import { validateId, validateObject } from './validator';
import { Validator, ValidatorHelper } from './validator.types';

/**
 * apollo
 */
const apolloHelper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new ApolloCustomError({
    message: error.message,
    code: 'BAD_REQUEST',
    extensions: params,
  });
};

export const apolloValidator: Validator = {
  validateId: (id) => validateId(id)(apolloHelper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(apolloHelper),
};

/**
 * fastify
 */
const fastifyHelper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new FastifyCustomError({
    message: `params: ${JSON.stringify(params)}\n${error.message}`,
    name: 'BadRequestError',
    statusCode: 400,
  });
};

export const fastifyValidator: Validator = {
  validateId: (id) => validateId(id)(fastifyHelper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(fastifyHelper),
};
