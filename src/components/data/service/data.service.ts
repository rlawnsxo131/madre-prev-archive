import { getConnection } from 'typeorm';
import DataQueryRepository from '../repository/data.query.repository';

function getData(id: number) {
  const connection = getConnection('default');
  return connection.getCustomRepository(DataQueryRepository).findOneById(id);
}

export default {
  getData,
};
