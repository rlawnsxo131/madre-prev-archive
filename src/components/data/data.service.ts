import { getCustomRepository } from 'typeorm';
import { dataError, DataRepository } from '.';

async function getDataById(id: number) {
  const dataRepo = getCustomRepository(DataRepository);
  const data = await dataRepo.findOne({ id });
  if (!data) {
    throw dataError.notFoundData;
  }
  return data;
}

export default {
  getDataById,
};
