const path = require('path');
module.exports = {
  entry: {
    activator_activate: './public/src/js/activator_activate.js',
  },
  output: {
    path: path.resolve(__dirname, 'public/js'),
    publicPath: '/assets/',
    filename: '[name].js',
  },
  module: {
    loaders: [
    { test: /\.jsx$/, exclude: /node_modules/, loader: 'babel-loader'},
    { test: /\.js$/, exclude: /node_modules/, loader: 'babel-loader'},
    { test: /\.less$/, loader: 'style!css!less'},
    ]
  }
};
