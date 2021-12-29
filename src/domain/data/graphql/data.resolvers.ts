import { IResolvers } from '@graphql-tools/utils';
import DataService from '../service/data.service';
import ApolloCustomError from '../../../lib/ApolloCustomError';
import { ApolloValidator } from '../../../lib/Validator';
import { GetDataParams, CreateDataParams } from '../interface/data.interface';
import { boolean, object, string } from 'joi';

const dataResolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
      ApolloValidator.validateId(id);
      const data = await DataService.findOne(id);
      if (!data) {
        throw new ApolloCustomError({
          message: 'Not Found Data',
          code: 'NOT_FOUND',
          extensions: { id },
        });
      }
      return data;
    },
  },
  Mutation: {
    async createData(_, params: CreateDataParams) {
      ApolloValidator.validateObject(
        object<CreateDataParams>({
          user_id: string().guid().required(),
          file_url: string().uri().required(),
          title: string().min(1).required(),
          is_public: boolean().required(),
          description: string().min(1).optional(),
        }),
        params,
      );
      const data = await DataService.create(params);
      return data;
    },
  },
};

export default dataResolvers;
