import 'jest';
import { fastifyErrorService } from '..';

describe('fastifyErrorService Test', () => {
  test('throwErrorValidation', () => {
    const name = 'NotFoundError';
    const message = 'Not found User';
    const statusCode = 404;
    try {
      fastifyErrorService.throwError({
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
