const webpack = require('webpack')

module.exports = {
    mode: 'production',
    devtool: 'source-map',
    plugins: [

        new webpack.DefinePlugin({
            'process.env.isLocal' : JSON.stringify(false),
            'process.env.back_end_url' : JSON.stringify('http://localhost:8000'),
        }),
    ],
}