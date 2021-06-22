const path = require('path');
const paths = require('./paths');
const { config } = require('dotenv');

function envConfig() {
  // const baseEnv = { REACT_APP_SSR: target === 'web' ? 'disabled' : 'enabled' };
  const { NODE_ENV } = process.env;
  config({
    path: path.resolve(paths.rootPath, `.env.${NODE_ENV}`),
  });
  return {
    'process.env': JSON.stringify(
      Object.keys(process.env)
        .filter((key) => /^REACT_APP/i.test(key))
        .reduce((env, key) => {
          env[key] = process.env[key];
          return env;
        }, {}),
    ),
  };
}

module.exports = envConfig;
