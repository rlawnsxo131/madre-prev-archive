import { userService } from '..';
import { errorCode, errorService } from '../../error';

async function getUserAction(id: number) {
  const user = await userService.getUserById(id);
  if (!user) {
    throw errorService.createApolloError({
      message: 'Not Found User',
      errorCode: errorCode.BAD_REQUEST,
      params: { id },
    });
  }
  return user;
}

export default {
  getUserAction,
};
