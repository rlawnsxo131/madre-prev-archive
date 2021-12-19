import { gql } from 'apollo-server-core';

export default gql`
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
    user(id: Int!): User
  }
`;
