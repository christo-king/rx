module.exports = {
    "context": __dirname,
    "entry": "./src/jsx/index.js",
    "output": {
        "path": __dirname + "/src/static/",
        "filename": "bundle.js"
    },
    module: {
        loaders: [
            {
                test: /.js?$/,
                loader: 'babel-loader',
                exclude: /node_modules/,
                query: {
                    presets: ['es2015', 'react']
                }
            },
            {
                test: /.scss$/,
                exclude: /node_modules/,
                loaders: ["style-loader", "css-loader", "sass-loader"]
            }
        ]
    },
};
