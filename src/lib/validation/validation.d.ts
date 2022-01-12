import Joi from 'joi';

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
