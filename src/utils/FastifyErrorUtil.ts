export namespace FastifyErrorUtil {
  interface FastifyCustomErrorParams {
    message: string;
    name:
      | 'BadRequestError'
      | 'NotFoundError'
      | 'InternalServerError'
      | 'UnauthorizedError'
      | 'ForbiddenError';
    statusCode: number;
  }

  class FastifyCustomError extends Error {
    statusCode: number;
    name: string;
    constructor({ message, name, statusCode }: FastifyCustomErrorParams) {
      super(message);
      this.name = name;
      this.statusCode = statusCode;
    }
  }

  export function throwError({
    message,
    name,
    statusCode,
  }: FastifyCustomErrorParams) {
    throw new FastifyCustomError({
      message,
      name,
      statusCode,
    });
  }
}
