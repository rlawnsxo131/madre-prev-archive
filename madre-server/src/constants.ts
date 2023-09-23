export const isProduction = process.env.NODE_ENV === 'production';
export const environmentFilename = `.env.${process.env.NODE_ENV}`;
