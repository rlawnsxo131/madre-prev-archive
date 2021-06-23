const path = require('path');
const paths = require('./paths');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const { WebpackManifestPlugin } = require('webpack-manifest-plugin');
const PnpWebpackPlugin = require('pnp-webpack-plugin');
const envConfig = require('./envConfig');

module.exports = () => {
  const clientEnv = envConfig();
  return {
    mode: 'development',
    entry: paths.entryPath,
    output: {
      path: paths.buildPath,
      publicPath: '/',
      filename: 'static/js/[name].[contenthash:8].js',
      chunkFilename: 'static/js/[name].[contenthash:8].js',
    },
    target: 'web',
    devtool: 'cheap-module-source-map',
    module: {
      rules: [
        {
          test: /\.(js|jsx|ts|tsx)$/,
          exclude: /node_modules/,
          use: [
            'babel-loader',
            {
              loader: 'ts-loader',
              options: {
                transpileOnly: true,
              },
            },
          ],
        },
        {
          test: [/\.bmp$/, /\.gif$/, /\.jpe?g$/, /\.png$/],
          oneOf: [
            {
              loader: 'url-loader',
              options: {
                name: 'static/media/[name].[contenthash:8].[ext]',
                limit: 10000,
              },
            },
            {
              loader: 'file-loader',
              options: {
                name: 'static/media/[name].[contenthash:8].[ext]',
                esModule: false,
              },
            },
          ],
        },
        {
          test: /\.css$/,
          use: [
            MiniCssExtractPlugin.loader,
            {
              loader: 'css-loader',
              options: {
                sourceMap: true,
              },
            },
          ],
        },
      ],
    },
    resolve: {
      modules: ['node_modules'],
      extensions: ['.tsx', '.ts', '.jsx', '.js'],
      fallback: {
        path: false,
      },
      plugins: [PnpWebpackPlugin],
    },
    resolveLoader: {
      plugins: [PnpWebpackPlugin.moduleLoader(module)],
    },
    optimization: {
      minimize: false,
      splitChunks: {
        chunks: 'all',
        name: false,
      },
      runtimeChunk: {
        name: (entrypoint) => `runtime-${entrypoint.name}`,
      },
    },
    plugins: [
      new HtmlWebpackPlugin({
        filename: 'index.html',
        template: path.resolve(paths.publicPath, 'index.html'),
        templateParameters: {
          env: {
            REACT_APP_PUBLIC_URL: '',
            REACT_APP_IMAGE_URL: path.resolve(__dirname, '../static'),
          },
        },
        filename: 'index.html',
      }),
      new MiniCssExtractPlugin({
        filename: 'static/css/[name].[contenthash:8].css',
        chunkFilename: 'static/css/[name].[contenthash:8].chunk.css',
      }),
      new WebpackManifestPlugin({
        fileName: 'asset-manifest.json',
        publicPath: '/',
        generate: (seed, files, entrypoints) => {
          const manifestFiles = files.reduce((manifest, file) => {
            manifest[file.name] = file.path;
            return manifest;
          }, seed);
          const entrypointFiles = entrypoints.main.filter(
            (fileName) => !fileName.endsWith('.map'),
          );
          return {
            files: manifestFiles,
            entrypoints: entrypointFiles,
          };
        },
      }),
      new webpack.DefinePlugin(clientEnv),
      new webpack.HotModuleReplacementPlugin(),
    ].filter(Boolean),
    cache: {
      type: 'memory',
    },
    devServer: {
      // host: '',
      publicPath: '/',
      port: 8080,
      open: true,
      hot: true,
      overlay: true,
      historyApiFallback: true,
      stats: 'errors-warnings',
    },
    stats: {
      builtAt: true,
      children: true,
      entrypoints: true,
      hash: true,
      modules: true,
      version: true,
      publicPath: true,
      // excludeAssets: [/\.(map|txt|html|jpg|png)$/, /\.json$/],
    },
  };
};
