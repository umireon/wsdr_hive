const path = require('path');
module.exports = {
    entry: './public/src/js/main.jsx',
    output: {
        path: path.resolve(__dirname, 'public/js'),
        publicPath: '/assets/',
        filename: 'main.js',
    },
    module: {
        loaders: [
            { test: /\.jsx$/, exclude: /node_modules/, loader: 'babel-loader'}
        ]
    }
};
