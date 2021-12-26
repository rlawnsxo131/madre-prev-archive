import Joi from 'joi';
import { ApolloErrorUtil } from '../../../utils';

export namespace UserValidationService {
  export function getUserParamsValidation(id: string) {
    const schema = Joi.string().guid().required();
    const { error } = schema.validate(id);
    if (!error) return;
    ApolloErrorUtil.throwError({
      message: error.message,
      code: 'BAD_REQUEST',
      params: { id },
    });
  }
}
