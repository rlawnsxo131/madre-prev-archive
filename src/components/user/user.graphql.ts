import { gql } from 'apollo-server-core';
import { IResolvers } from '@graphql-tools/utils';
import { userActionService } from '.';

const typeDef = gql`
  type User {
    id: Int!
    name: String!
  }
  extend type Query {
    user(id: Int!): User
  }
`;

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, args) {
      const { id } = args;
      const user = await userActionService.getUserAction(id);
      return user;
    },
  },
  Mutation: {},
};

export default {
  typeDef,
  resolvers,
};
