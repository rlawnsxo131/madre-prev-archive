import 'jest';
import { apolloErrorService } from '..';

describe('apolloErrorService Test', () => {
  test('throwErrorValidation: type NOT_FOUND and id to 1', () => {
    const id = 1;
    const code = 'NOT_FOUND';
    try {
      apolloErrorService.throwApolloErrorValidation({
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
});
