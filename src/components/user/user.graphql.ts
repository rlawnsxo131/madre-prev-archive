import { gql } from 'apollo-server-core';
import { IResolvers } from '@graphql-tools/utils';
import { UserService } from '.';

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
      const user = await UserService.findById(id);
      return user;
    },
  },
  Mutation: {},
};

const UserGraphQL = {
  typeDef,
  resolvers,
};

export default UserGraphQL;
