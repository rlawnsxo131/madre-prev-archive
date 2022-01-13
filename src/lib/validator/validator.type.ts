import Joi from 'joi';

export type ValidatorHelper = (
  error: Joi.ValidationError | undefined,
  params: Record<string, any>,
) => void;

export type ValidateId = (id: string) => (helper: ValidatorHelper) => void;
export type ValidateObject = (
  schema: Joi.Schema,
  parmas: Record<string, any>,
) => (helper: ValidatorHelper) => void;

export interface Validator {
  validateId: (id: string) => ReturnType<ReturnType<ValidateId>>;
  validateObject: (
    schema: Joi.Schema,
    parmas: Record<string, any>,
  ) => ReturnType<ReturnType<ValidateObject>>;
}
