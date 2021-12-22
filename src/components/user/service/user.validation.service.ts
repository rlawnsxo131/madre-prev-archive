import Joi from 'joi';
import { UserInputError } from 'apollo-server-core';

function getUserParamsValidation(id: string) {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  if (!error) return;
  throw new UserInputError(error.message, { id });
}

export default {
  getUserParamsValidation,
};
