import { GraphQLError } from 'graphql';

export default function apolloFormatError(error: GraphQLError) {
  console.error(
    '------------------------------- ERROR INFO -------------------------------',
  );
  console.error(error.toJSON());
  console.error(error.extensions.exception?.stacktrace);
  console.error(
    '------------------------------- ERROR INFO -------------------------------',
  );
  return error;
}
