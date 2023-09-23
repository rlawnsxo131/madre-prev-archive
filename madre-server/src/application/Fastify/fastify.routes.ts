import { FastifyPluginCallback } from 'fastify';
import { authRoute } from '../../domain/auth';

const fastifyRoutes: FastifyPluginCallback = (fastify, opts, done) => {
  fastify.register(authRoute, { prefix: '/auth' });
  done();
};

export default fastifyRoutes;
