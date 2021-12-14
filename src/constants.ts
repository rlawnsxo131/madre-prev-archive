export const isProduction = process.env.NODE_ENV === 'production';
export const environmentFilename = `.env.${process.env.NODE_ENV}`;
export const errorCode = {
  NOT_FOUND: 'NOT_FOUND',
  FORBIDDEN: 'FORBIDDEN',
  UNAUTHENTICATED: 'UNAUTHENTICATED',
  INTERNAL_SERVER_ERROR: 'INTERNAL_SERVER_ERROR',
  BAD_REQUEST: 'BAD_REQUEST',
  BAD_USER_INPUT: 'BAD_USER_INPUT',
};
