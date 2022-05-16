const webpack = require('webpack')
const ReactRefreshWebpackPlugin = require('@pmmmwh/react-refresh-webpack-plugin');

module.exports = {
    mode: 'development',
    devServer: {
        hot: true,
        open: true,
    },
    devtool: 'cheap-module-source-map',
    plugins: [
        new webpack.DefinePlugin({
            'process.env.isLocal' : JSON.stringify(true),
            'process.env.back_end_url' : JSON.stringify('http://localhost:8000')
        }),
        new ReactRefreshWebpackPlugin(),
    ],
}
