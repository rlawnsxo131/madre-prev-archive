import { IResolvers } from '@graphql-tools/utils';
import { DataService, DataValidationService } from '..';
import { ApolloErrorUtil } from '../../../utils';
import { CreateDataParams, GetDataParams } from '../interface/data.interface';

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
      DataValidationService.getDataParamsValidation(id);
      const data = await DataService.getData(id);
      if (!data) {
        ApolloErrorUtil.throwError({
          message: 'Not Found Data',
          code: 'NOT_FOUND',
          params: { id },
        });
      }
      return data;
    },
  },
  Mutation: {
    async createData(_, params: CreateDataParams) {
      DataValidationService.createDataParamsValidation(params);
      const data = await DataService.createData(params);
      return data;
    },
  },
};

export default resolvers;
