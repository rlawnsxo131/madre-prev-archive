const path = require('path');
const paths = require('./paths');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');
const { WebpackManifestPlugin } = require('webpack-manifest-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const PnpWebpackPlugin = require('pnp-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin');

// client environment
const initializeEnvironment = require('./initializeEnvironment');
const clientEnvironment = initializeEnvironment();
const { NODE_ENV, REACT_APP_PUBLIC_URL, REACT_APP_IMAGE_URL } = process.env;

const isProduction = NODE_ENV === 'production';
const prefix = isProduction ? '/' : '';
const publicPath = `${REACT_APP_PUBLIC_URL}${prefix}`;
const reactAppPublicUrl = isProduction ? REACT_APP_PUBLIC_URL : prefix;

module.exports = {
  mode: 'production',
  entry: paths.entryPath,
  output: {
    path: paths.buildPath,
    publicPath,
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
    minimize: true,
    minimizer: [new CssMinimizerPlugin(), new TerserPlugin()],
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
          REACT_APP_PUBLIC_URL: reactAppPublicUrl,
          REACT_APP_IMAGE_URL: REACT_APP_IMAGE_URL,
        },
      },
      minify: {
        removeComments: true,
        collapseWhitespace: true,
        removeRedundantAttributes: true,
        useShortDoctype: true,
        removeEmptyAttributes: true,
        removeStyleLinkTypeAttributes: true,
        keepClosingSlash: true,
        minifyJS: true,
        minifyCSS: true,
        minifyURLs: true,
      },
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
    new webpack.DefinePlugin(clientEnvironment),
    new ForkTsCheckerWebpackPlugin(),
    new CleanWebpackPlugin(),
  ].filter(Boolean),
  cache: {
    type: 'filesystem',
    buildDependencies: {
      config: [__filename],
    },
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
