type ErrorName =
  | 'BadRequestError'
  | 'NotFoundError'
  | 'InternalServerError'
  | 'UnauthorizedError'
  | 'ForbiddenError';

interface CustomErrorParams {
  message: string;
  name: ErrorName;
  statusCode: number;
}

class CustomError extends Error {
  statusCode: number;
  name: string;
  constructor({ message, name, statusCode }: CustomErrorParams) {
    super(message);
    this.name = name;
    this.statusCode = statusCode;
  }
}

interface ThrowErrorValidationParams {
  resolver: (params?: any) => boolean;
  message: string;
  name: ErrorName;
  statusCode: number;
}

function throwErrorValidation({
  resolver,
  message,
  name,
  statusCode,
}: ThrowErrorValidationParams) {
  if (!resolver()) return;
  throw new CustomError({
    message,
    name,
    statusCode,
  });
}

export default {
  throwErrorValidation,
};
