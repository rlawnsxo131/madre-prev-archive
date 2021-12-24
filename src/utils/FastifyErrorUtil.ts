export namespace FastifyErrorUtil {
  type ErrorName =
    | 'BadRequestError'
    | 'NotFoundError'
    | 'InternalServerError'
    | 'UnauthorizedError'
    | 'ForbiddenError';

  interface FastifyCustomErrorParams {
    message: string;
    name: ErrorName;
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
