var path = require('path');

var PROD = JSON.parse(process.env.PROD_ENV || '0');
var BUILD_DIR = path.resolve(__dirname, 'src/static/');
var JSX_DIR = path.resolve(__dirname, 'src/jsx');
var glob = require("glob");

var config = {
    entry: glob.sync(JSX_DIR + '**/*.jsx'),
    output: {
        path: BUILD_DIR,
        filename: 'bundle.js'
    },
    module: {
        loaders: [
            {
                test: /\.jsx/,
                include: JSX_DIR,
                loader: 'babel-loader',
                query: {
                    presets: ['es2015', 'react']
                }
            }
        ]
    }
};

module.exports = config;

var path = require('path');
var webpack = require('webpack');
