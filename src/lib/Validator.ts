import Joi from 'joi';
import ApolloCustomError from './ApolloCustomError';
import FastifyCustomError from './FastifyCustomError';

type Helper = (
  error: Joi.ValidationError | undefined,
  params: Record<string, any>,
) => void;

const apolloHelper: Helper = (error, params) => {
  if (!error) return;
  throw new ApolloCustomError({
    message: error.message,
    code: 'BAD_REQUEST',
    extensions: params,
  });
};

const fastifyHelper: Helper = (error, params) => {
  if (!error) return;
  throw new FastifyCustomError({
    message: `params: ${JSON.stringify(params)}\n${error.message}`,
    name: 'BadRequestError',
    statusCode: 400,
  });
};

class Validator {
  helper: Helper;

  constructor(helper: Helper) {
    this.helper = helper;
  }

  validateId(id: string = '') {
    const schema = Joi.string().guid().required();
    const { error } = schema.validate(id);
    this.helper(error, { id });
  }

  validateObject(schema: Joi.Schema, params: Record<string, any> = {}) {
    const { error } = schema.validate(params);
    if (!error) return;
    this.helper(error, params);
  }
}

export const ApolloValidator = new Validator(apolloHelper);
export const FastifyValidator = new Validator(fastifyHelper);
