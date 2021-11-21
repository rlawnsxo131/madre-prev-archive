export default {
  isProduction: process.env.NODE_ENV === 'production',
  environmentFilename: `.env.${process.env.NODE_ENV}`,
};
