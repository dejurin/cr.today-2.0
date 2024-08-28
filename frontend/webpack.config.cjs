const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');

module.exports = {
  entry: {
    common: './src/index.ts',
  },
  output: {
    path: path.resolve(__dirname, '../webapp/static'),
    filename: '[name].bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.css$/i,
        use: [MiniCssExtractPlugin.loader, 'css-loader', 'postcss-loader'],
      },
    ],
  },
  optimization: {
    minimize: true,
    minimizer: [
      '...', // Extends existing minimizers (i.e., `terser-webpack-plugin`)
      new CssMinimizerPlugin(),
    ],
  },
  devServer: {
    static: {
      directory: path.join(__dirname, '../static'),
    },
    compress: true,
    port: 9000,
    hot: true,
    liveReload: true,
    open: true,
    historyApiFallback: true,
  },
  resolve: {
    extensions: ['.ts', '.js'],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name].bundle.css',
    }),
  ],
};