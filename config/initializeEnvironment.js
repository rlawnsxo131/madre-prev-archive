const path = require('path');
const paths = require('./paths');
const { config } = require('dotenv');

function initializeEnvironment() {
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

module.exports = initializeEnvironment;
