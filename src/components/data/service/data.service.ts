import { ApolloError } from 'apollo-server-core';
import { getCustomRepository } from 'typeorm';
import { errorCode } from '../../../constants';
import DataQueryRepository from '../repository/data.query.repository';

async function getDataById(id: number) {
  const dataQueryRepository = getCustomRepository(DataQueryRepository);
  const data = await dataQueryRepository.findOneById(id);
  if (!data) {
    throw new ApolloError('Not Found Data', errorCode.NOT_FOUND, { id });
  }
  return data;
}

export default {
  getDataById,
};
