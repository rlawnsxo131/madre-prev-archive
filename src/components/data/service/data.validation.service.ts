import Joi from 'joi';
import { CreateDataParams } from '../interface/data.interface';
import { UserInputError } from 'apollo-server-core';

function getDataParamsValidation(id: string) {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  if (!error) return;
  throw new UserInputError(error.message, { id });
}

function createDataParamsValidation(params: CreateDataParams) {
  const schema = Joi.object<CreateDataParams>({
    user_id: Joi.string().guid().required(),
    file_url: Joi.string().uri().required(),
    title: Joi.string().min(1).required(),
    is_public: Joi.boolean().required(),
    description: Joi.string().min(1).optional(),
  });
  const { error } = schema.validate(params);
  if (!error) return;
  throw new UserInputError(error.message, params);
}

export default {
  getDataParamsValidation,
  createDataParamsValidation,
};
