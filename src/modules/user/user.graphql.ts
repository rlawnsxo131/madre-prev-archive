import { gql } from 'apollo-server-core';
import { IResolvers } from '@graphql-tools/utils';
import { UserService } from '.';

export const typeDef = gql`
  type User {
    id: Int!
    name: String!
  }
  extend type Query {
    user(id: Int!): User
  }
`;

interface UserArgs {
  id: number;
}

export const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(parent, args, context, info) {
      const { id } = args as UserArgs;
      const user = await UserService.findOne(id);
      return user;
    },
  },
  Mutation: {},
};
