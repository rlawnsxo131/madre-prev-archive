import Joi from 'joi';
import { ApolloErrorUtil } from '../../../utils/ApolloErrorUtil';
import { CreateDataParams } from '../interface/data.interface';

export namespace DataValidationService {
  export function getDataParamsValidation(id: string) {
    const schema = Joi.string().guid().required();
    const { error } = schema.validate(id);
    if (!error) return;
    ApolloErrorUtil.throwError({
      message: error.message,
      code: 'BAD_REQUEST',
      params: { id },
    });
  }

  export function createDataParamsValidation(
    params: CreateDataParams | Record<string, any> = {},
  ) {
    const schema = Joi.object<CreateDataParams>({
      user_id: Joi.string().guid().required(),
      file_url: Joi.string().uri().required(),
      title: Joi.string().min(1).required(),
      is_public: Joi.boolean().required(),
      description: Joi.string().min(1).optional(),
    });
    const { error } = schema.validate(params);
    if (!error) return;
    ApolloErrorUtil.throwError({
      message: error.message,
      code: 'BAD_USER_INPUT',
      params,
    });
  }
}
