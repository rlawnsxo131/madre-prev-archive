import { FastifyPluginCallback } from 'fastify';
import fp from 'fastify-plugin';

const callback: FastifyPluginCallback = (fastify, opts, done) => {
  fastify.decorateRequest('user', null);
  fastify.addHook('onRequest', (request, reply) => {
    // console.log(request.cookies);
  });

  done();
};

const jwtPlugin = fp(callback, {
  name: 'jwtPlugin',
});

export default jwtPlugin;
