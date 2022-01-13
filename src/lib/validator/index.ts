import Joi from 'joi';
import ApolloCustomError from '../ApolloCustomError';
import FastifyCustomError from '../FastifyCustomError';

/**
 * type definition
 */
type ValidatorHelper = (
  error: Joi.ValidationError | undefined,
  params: Record<string, any>,
) => void;

type ValidateId = (id: string) => (helper: ValidatorHelper) => void;
type ValidateObject = (
  schema: Joi.Schema,
  parmas: Record<string, any>,
) => (helper: ValidatorHelper) => void;

interface Validator {
  validateId: (id: string) => ReturnType<ReturnType<ValidateId>>;
  validateObject: (
    schema: Joi.Schema,
    parmas: Record<string, any>,
  ) => ReturnType<ReturnType<ValidateObject>>;
}

/**
 * common function
 */
const validateId: ValidateId = (id) => (helper) => {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  helper(error, { id });
};

const validateObject: ValidateObject = (schema, params) => (helper) => {
  const { error } = schema.validate(params);
  helper(error, params);
};

/**
 * helper
 */
const apolloHelper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new ApolloCustomError({
    message: error.message,
    code: 'BAD_REQUEST',
    extensions: params,
  });
};

const fastifyHelper: ValidatorHelper = (error, params) => {
  if (!error) return;
  throw new FastifyCustomError({
    message: `params: ${JSON.stringify(params)}\n${error.message}`,
    name: 'BadRequestError',
    statusCode: 400,
  });
};

/**
 * export validator
 */
export const apolloValidator: Validator = {
  validateId: (id) => validateId(id)(apolloHelper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(apolloHelper),
};

export const fastifyValidator: Validator = {
  validateId: (id) => validateId(id)(fastifyHelper),
  validateObject: (schema, params = {}) =>
    validateObject(schema, params)(fastifyHelper),
};
