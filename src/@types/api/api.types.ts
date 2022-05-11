/**
 * rtk catch block error response data set
 * {
 *   error: {
 *     status: 500,
 *     data: {
 *       message: "InternalServerError"
 *       status: 500
 *     }
 *   }
 *   isUnhandledError: false
 *   meta: {request: Request, response: Response}
 * }
 */

type ResponseErrors =
  | {
      message: 'BadRequest';
      status: 400;
    }
  | {
      message: 'NotFound';
      status: 404;
    }
  | {
      message: 'InternalServerError';
      status: 500;
    }
  | {
      message: 'Unauthorized';
      status: 401;
    }
  | {
      message: 'Forbidden';
      status: 403;
    }
  | {
      message: 'Conflict';
      status: 409;
    };

export interface ResponseError {
  error: {
    status: number;
    data: ResponseErrors;
  };
  isUnhandledError: boolean;
  meta: {
    request: Request;
    response: Response;
  };
}
