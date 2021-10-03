import { FastifyPluginCallback } from 'fastify';

const routes: FastifyPluginCallback = (fastify, opts, done) => {
  fastify.get('/health', async (request, reply) => {
    reply.status(200).send({ hello: 'world' });
  });
  done();
};

export default routes;
