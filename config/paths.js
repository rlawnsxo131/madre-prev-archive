const path = require('path');
const rootPath = path.resolve(__dirname, '../');
const entryPath = path.resolve(__dirname, '../src', 'index.tsx');
const staticPath = path.resolve(__dirname, '../src', 'static');
const publicPath = path.resolve(__dirname, '../public');
const buildPath = path.resolve(__dirname, '../dist');

module.exports = {
  rootPath,
  entryPath,
  staticPath,
  publicPath,
  buildPath,
};
