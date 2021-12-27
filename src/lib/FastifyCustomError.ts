type FastifyErrorName =
  | 'BadRequestError'
  | 'NotFoundError'
  | 'InternalServerError'
  | 'UnauthorizedError'
  | 'ForbiddenError';

interface FastifyCustomErrorParams {
  message: string;
  name: FastifyErrorName;
  statusCode: number;
}

export default class FastifyCustomError extends Error {
  statusCode: number;
  name: string;
  constructor({ message, name, statusCode }: FastifyCustomErrorParams) {
    super(message);
    this.name = name;
    this.statusCode = statusCode;
  }
}
