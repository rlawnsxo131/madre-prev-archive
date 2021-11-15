import { dataService } from '..';
import { errorCode, errorService } from '../../error';

async function getDataAction(id: number) {
  const data = await dataService.getData(id);
  if (!data) {
    throw errorService.createApolloError({
      message: 'Not Found Data',
      errorCode: errorCode.NOT_FOUND,
      params: { id },
    });
  }
  return data;
}

export default {
  getDataAction,
};
