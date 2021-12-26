import Joi from 'joi';
import { apolloErrorManager } from '../../../lib';

function getUserParamsValidation(id: string) {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  if (!error) return;
  apolloErrorManager.throwError({
    message: error.message,
    code: 'BAD_REQUEST',
    params: { id },
  });
}

export default {
  getUserParamsValidation,
};
