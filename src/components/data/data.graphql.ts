import { IResolvers } from '@graphql-tools/utils';
import { gql } from 'apollo-server-core';
import { dataService } from '.';

const typeDef = gql`
  type Data {
    id: Int!
    data: [Int]
  }
  extend type Query {
    data(id: Int!): Data
  }
`;

const resolvers: IResolvers = {
  Query: {
    async data(_, args) {
      const { id } = args;
      const data = await dataService.findById(id);
      return data;
    },
  },
};

export default {
  typeDef,
  resolvers,
};
