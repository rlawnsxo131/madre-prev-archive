import { IResolvers } from '@graphql-tools/utils';
import { gql } from 'apollo-server-core';
import { DataService } from '.';

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
      const data = await DataService.findById(id);
      return data;
    },
  },
};

const DataGraphQL = {
  typeDef,
  resolvers,
};

export default DataGraphQL;
