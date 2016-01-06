module.exports = function(grunt) {
    var webpackConfig = require('./webpack.config.js');
    grunt.loadNpmTasks('grunt-webpack');
    grunt.initConfig({
        'webpack-dev-server': {
            options: {
                webpack: webpackConfig,
                publicPath: "/"
            },
            start: {
                keepAlive: true,
                publicPath: webpackConfig.output.publicPath,
                contentBase: "./public",
                proxy: {
                    "/api/*": {
                        target: "http://localhost:9000",
                        ws: true
                    }
                }
            }
        }
    });
    grunt.registerTask('default', ['webpack-dev-server:start']);
};
