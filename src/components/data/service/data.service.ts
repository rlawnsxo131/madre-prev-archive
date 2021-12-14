import { ApolloError } from 'apollo-server-core';
import { getCustomRepository } from 'typeorm';
import DataQueryRepository from '../repository/data.query.repository';
import { ERROR_CODE } from '../../../constants';

async function getDataById(id: number) {
  const dataQueryRepository = getCustomRepository(DataQueryRepository);
  const data = await dataQueryRepository.findOneById(id);
  if (!data) {
    throw new ApolloError('Not Found Data', ERROR_CODE.NOT_FOUND, { id });
  }
  return data;
}

export default {
  getDataById,
};
