import { ApolloServer } from 'apollo-server-fastify';
import {
  ApolloServerPluginDrainHttpServer,
  ApolloServerPluginLandingPageDisabled,
  ApolloServerPluginLandingPageGraphQLPlayground,
  GraphQLRequestContext,
} from 'apollo-server-core';
import { FastifyInstance } from 'fastify';
import apolloSchema from './apollo.schema';
import { isProduction } from '../../constants';
import { GraphQLError, print } from 'graphql';

export default class Apollo {
  private readonly app: ApolloServer;

  constructor(fastify: FastifyInstance) {
    this.app = new ApolloServer({
      schema: apolloSchema,
      context: ({ request, reply }) => {
        request.log.info(request.id);
      },
      formatError: this.formatError,
      plugins: [
        this.fastifyAppClosePlugin(fastify),
        ApolloServerPluginDrainHttpServer({ httpServer: fastify.server }),
        {
          async requestDidStart(_) {
            console.log('GraphQL requestDidStart');
            return {
              async parsingDidStart(_) {
                console.log('GraphQL parsingDidStart');
              },
              async validationDidStart(requestContext) {
                console.log(
                  'GraphQL validationDidStart: ',
                  requestContext.document.loc?.source,
                );
              },
            };
          },
        },
        isProduction
          ? ApolloServerPluginLandingPageDisabled()
          : ApolloServerPluginLandingPageGraphQLPlayground(),
      ],
      debug: !isProduction,
    });
  }

  private formatError(error: GraphQLError) {
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

  private fastifyAppClosePlugin(fastify: FastifyInstance) {
    return {
      async serverWillStart() {
        return {
          async drainServer() {
            await fastify.close();
          },
        };
      },
    };
  }

  getApp() {
    return this.app;
  }

  start() {
    return this.app.start();
  }

  createHandler() {
    return this.app.createHandler({ cors: false });
  }
}
