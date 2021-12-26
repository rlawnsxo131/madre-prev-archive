import 'jest';
import { dataValidationService } from '..';
import { CreateDataParams } from '../interface/data.interface';

describe('dataValidationService Test', () => {
  test('getDataParamsValidation', () => {
    const id = '';
    try {
      dataValidationService.getDataParamsValidation(id);
    } catch (e: any) {
      expect(e.extensions.code).toBe('BAD_REQUEST');
      expect(e.extensions.id).toBe(id);
    }
  });

  test('createDataParamsValidation', () => {
    const createDataParams: CreateDataParams = {
      user_id: '',
      file_url: '',
      is_public: false,
      title: '',
      description: '',
    };
    try {
      dataValidationService.createDataParamsValidation(createDataParams);
    } catch (e: any) {
      expect(e.extensions.code).toBe('BAD_USER_INPUT');
      expect(e.extensions.user_id).toBe(createDataParams.user_id);
      expect(e.extensions.file_url).toBe(createDataParams.file_url);
      expect(e.extensions.is_public).toBe(createDataParams.is_public);
      expect(e.extensions.title).toBe(createDataParams.title);
      expect(e.extensions.description).toBe(createDataParams.description);
    }
  });
});
