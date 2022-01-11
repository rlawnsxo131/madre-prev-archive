import Joi from 'joi';

export type ValidatorHelper = (
  error: Joi.ValidationError | undefined,
  params: Record<string, any>,
) => void;

export interface Validator {
  validateId: (id: string) => ReturnType<ReturnType<typeof validateId>>;
  validateObject: (
    schema: Joi.Schema,
    parmas: Record<string, any>,
  ) => ReturnType<ReturnType<typeof validateObject>>;
}

export const validateId = (id: string) => (helper: ValidatorHelper) => {
  const schema = Joi.string().guid().required();
  const { error } = schema.validate(id);
  helper(error, { id });
};

export const validateObject =
  (schema: Joi.Schema, params: Record<string, any>) =>
  (helper: ValidatorHelper) => {
    const { error } = schema.validate(params);
    helper(error, params);
  };
