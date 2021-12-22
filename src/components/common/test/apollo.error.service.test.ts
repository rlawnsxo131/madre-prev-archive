import 'jest';
import { apolloErrorService, fastifyErrorService } from '..';

describe('apolloErrorService Test', () => {
  test('throwApolloErrorService.throwErrorValidation: type NOT_FOUND and id to 1', () => {
    const id = 1;
    const code = 'NOT_FOUND';
    try {
      apolloErrorService.throwErrorValidation({
        resolver: () => true,
        message: 'Not Found Data',
        code,
        params: { id },
      });
    } catch (e: any) {
      expect(e.extensions.code).toBe(code);
      expect(e.extensions.id).toBe(id);
    }
  });

  test('throwFastifyErrorService.throwErrorValidation', () => {
    const name = 'NotFoundError';
    const message = 'Not found User';
    const statusCode = 404;
    try {
      fastifyErrorService.throwErrorValidation({
        resolver: () => true,
        message,
        name,
        statusCode,
      });
    } catch (e: any) {
      expect(e.name).toBe(name);
      expect(e.message).toBe(message);
      expect(e.statusCode).toBe(statusCode);
    }
  });
});
