import { gql } from 'apollo-server-core';

const typeDef = gql`
  type User {
    id: ID!
    email: String!
    username: String
    display_name: String!
    photo_url: String
    created_at: Date!
    updated_at: Date!
  }
  extend type Query {
    user(id: ID!): User
  }
`;

export default typeDef;
