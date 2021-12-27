import Joi from 'joi';
import ApolloCustomError from '../../../lib/ApolloCustomError';

export function getUserParamsValidation(id: string) {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  if (!error) return;
  throw new ApolloCustomError({
    message: error.message,
    code: 'BAD_REQUEST',
    extensions: { id },
  });
}
