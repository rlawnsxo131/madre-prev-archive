import { IResolvers } from '@graphql-tools/utils';
import { dataService, dataValidationService } from '..';
import { errorService } from '../../error';
import { CreateDataParams, GetDataParams } from '../interface/data.interface';

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
      const validation = dataValidationService.getDataParamsValidation(id);
      errorService.throwApolloError({
        resolver: () => !validation,
        message: 'data query validation error',
        code: 'BAD_USER_INPUT',
        params: { id },
      });
      const data = await dataService.getData(id);
      errorService.throwApolloError({
        resolver: () => !data,
        message: 'Not Found Data',
        code: 'NOT_FOUND',
        params: { id },
      });
      return data;
    },
  },
  Mutation: {
    async createData(_, params: CreateDataParams) {
      const validation =
        dataValidationService.createDataParamsValidation(params);
      errorService.throwApolloError({
        resolver: () => !validation,
        message: 'createData mutation validation error',
        code: 'BAD_USER_INPUT',
        params,
      });
      const data = await dataService.createData(params);
      return data;
    },
  },
};

export default resolvers;
