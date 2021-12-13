import { getCustomRepository } from 'typeorm';
import { errorCode, errorService } from '../../error';
import DataQueryRepository from '../repository/data.query.repository';

async function getDataById(id: number) {
  const dataQueryRepository = getCustomRepository(DataQueryRepository);
  const data = await dataQueryRepository.findOneById(id);
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
  getDataById,
};
