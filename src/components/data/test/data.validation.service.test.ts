import 'jest';
import { ApolloError } from 'apollo-server-core';
import { dataValidationService } from '..';
import { errorService } from '../../error';
import { CreateDataParams } from '../interface/data.interface';

describe('dataValidationService Test', () => {
  test('createDataParamsValidation', async () => {
    const createDataParams: CreateDataParams = {
      user_id: '',
      file_url: '',
      is_public: false,
      title: '',
      description: '',
    };
    const code = 'BAD_USER_INPUT';
    try {
      const validation =
        dataValidationService.createDataParamsValidation(createDataParams);
      errorService.throwApolloError({
        resolver: () => !validation,
        message: 'createData mutation validation error',
        code,
        params: createDataParams,
      });
    } catch (e) {
      const error = e as ApolloError;
      expect(error.extensions.code).toBe(code);
      expect(error.extensions.user_id).toBe(createDataParams.user_id);
      expect(error.extensions.file_url).toBe(createDataParams.file_url);
      expect(error.extensions.is_public).toBe(createDataParams.is_public);
      expect(error.extensions.title).toBe(createDataParams.title);
      expect(error.extensions.description).toBe(createDataParams.description);
    }
  });
});
