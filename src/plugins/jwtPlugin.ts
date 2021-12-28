import { FastifyPluginCallback } from 'fastify';
import fp from 'fastify-plugin';

const callback: FastifyPluginCallback = async (fastify, opts, done) => {
  fastify.decorateRequest('user', null);
  fastify.addHook('onRequest', async (request, reply) => {
    // console.log(request.cookies);
  });

  done();
};

const jwtPlugin = fp(callback, {
  name: 'jwtPlugin',
});

export default jwtPlugin;
