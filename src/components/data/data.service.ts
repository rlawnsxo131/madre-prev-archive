import { getCustomRepository } from 'typeorm';
import { dataError, DataRepository } from '.';
import { errorCode, errorService } from '../error';

async function getDataById(id: number) {
  const dataRepo = getCustomRepository(DataRepository);
  const data = await dataRepo.findOne({ id });
  if (!data) {
    throw errorService.createError({
      message: dataError.errorMessage.NotFoundData,
      errorCode: errorCode.NOT_FOUND,
      params: { id },
    });
  }
  return data;
}

export default {
  getDataById,
};
