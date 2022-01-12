import Joi from 'joi';
import { ValidateId, ValidateObject } from './validation';

export const validateId: ValidateId = (id) => (helper) => {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  helper(error, { id });
};

export const validateObject: ValidateObject = (schema, params) => (helper) => {
  const { error } = schema.validate(params);
  helper(error, params);
};
