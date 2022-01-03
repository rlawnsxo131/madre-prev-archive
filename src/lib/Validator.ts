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
  private readonly helper: Helper;

  constructor(helper: Helper) {
    this.helper = helper;
  }

  validateId(id: string) {
    const schema = Joi.string().guid().required();
    const { error } = schema.validate(id);
    this.helper(error, { id });
  }

  validateObject(schema: Joi.Schema, params: Record<string, any> = {}) {
    const { error } = schema.validate(params);
    this.helper(error, params);
  }
}

const ApolloValidator = new Validator(apolloHelper);
const FastifyValidator = new Validator(fastifyHelper);

export { ApolloValidator, FastifyValidator };
