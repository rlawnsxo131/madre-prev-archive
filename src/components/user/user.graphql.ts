import { gql } from 'apollo-server-core';
import { IResolvers } from '@graphql-tools/utils';
import { userService } from '.';

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
      const user = await userService.findById(id);
      return user;
    },
  },
  Mutation: {},
};

export default {
  typeDef,
  resolvers,
};
