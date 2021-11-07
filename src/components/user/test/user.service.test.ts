import 'jest';
import { Connection } from 'typeorm';
import { userError, userService } from '..';
import { Database } from '../../../datastore';
import initializeEnvironment from '../../../lib/initializeEnvironment';
import { errorCode, errorService } from '../../error';

describe('userService Test', () => {
  let connection: Connection | null = null;

  beforeAll(async () => {
    initializeEnvironment();
    const database = new Database();
    connection = await database.getConnection();
  });

  afterAll(async () => {
    await connection?.close();
  });

  test('getUserById id to 0', async () => {
    const id = 0;
    try {
      await userService.getUserById(id);
    } catch (e) {
      expect(e).toStrictEqual(
        errorService.createError({
          message: userError.errorMessage.notFoundUser,
          errorCode: errorCode.NOT_FOUND,
          params: { id },
        }),
      );
    }
  });
});
