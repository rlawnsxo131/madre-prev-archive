import { FastifyPluginCallback } from 'fastify';
import authRoute from '../../domain/auth/route/auth.route';

const fastifyRoutes: FastifyPluginCallback = (fastify, opts, done) => {
  fastify.register(authRoute, { prefix: '/auth' });
  done();
};

export default fastifyRoutes;
