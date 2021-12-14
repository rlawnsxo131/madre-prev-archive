import { IResolvers } from '@graphql-tools/utils';
import { ApolloError } from 'apollo-server-core';
import { dataService } from '..';
import { ERROR_CODE } from '../../../constants';

interface DataArgs {
  id: number;
}

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: DataArgs) {
      const data = await dataService.getData(id);
      if (!data) {
        throw new ApolloError('Not Found Data', ERROR_CODE.NOT_FOUND, { id });
      }
      return data;
    },
  },
};

export default resolvers;
