import { gql } from 'apollo-server-core';

export default gql`
  type Data {
    id: Int!
    user_id: Int!
    file_url: String!
    title: String!
    description: String
    is_public: Boolean!
    created_at: Date!
    updated_at: Date!
  }
  extend type Query {
    data(id: Int!): Data!
  }
`;
