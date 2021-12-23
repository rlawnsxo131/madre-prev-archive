type FastifyCustomErrorName =
  | 'BadRequestError'
  | 'NotFoundError'
  | 'InternalServerError'
  | 'UnauthorizedError'
  | 'ForbiddenError';

interface FastifyCustomErrorParams {
  message: string;
  name: FastifyCustomErrorName;
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

function throwFastifyErrorValidation({
  resolver,
  message,
  name,
  statusCode,
}: FastifyCustomErrorParams & MadreResolveValidationFunctionInObject) {
  if (!resolver()) return;
  throw new FastifyCustomError({
    message,
    name,
    statusCode,
  });
}

export default {
  throwFastifyErrorValidation,
};
