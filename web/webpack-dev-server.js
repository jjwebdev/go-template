var webpack = require('webpack'),
		WebpackDevServer = require('webpack-dev-server'),
		config = require(process.cwd() + '/src/config'),
		webpackConfig = require('./webpack.config.js')

new WebpackDevServer(webpack(webpackConfig), {
	hot: true,
	inline: true,
	historyApiFallback: true,
	progress: true,
	colors: true,
	inline: true,
	watch: true,
	displayErrorDetails: true
}).listen(config.PORT, 'localhost', function (err) {
	if (err) {
		return console.log(err)
	}

	console.log('Listening at http://localhost:' + config.PORT)
})
