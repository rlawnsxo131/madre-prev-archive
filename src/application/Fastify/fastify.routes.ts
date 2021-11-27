import { FastifyPluginCallback } from 'fastify';
import { authRoute } from '../../components/auth';

const routes: FastifyPluginCallback = (fastify, opts, done) => {
  fastify.get('/health', (request, reply) => {
    reply.status(200).send({ hello: 'world' });
  });

  fastify.register(authRoute, { prefix: '/auth' });

  done();
};

export default routes;
