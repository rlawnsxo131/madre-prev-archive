import { getConnection } from 'typeorm';
import { DataQueryRepository } from '..';

function getData(id: string) {
  const connection = getConnection('default');
  return connection.getCustomRepository(DataQueryRepository).findOneById(id);
}

export default {
  getData,
};
