var webpack = require('webpack'),
		CleanWebpackPlugin = require('clean-webpack-plugin'),
		HtmlWebpackPlugin = require('html-webpack-plugin'),
		ENV = process.env.NODE_ENV || 'local',
		config = require(process.cwd() + '/src/config'),
		getHtmlPluginTemplates = require(process.cwd() + '/src/helpers/getHtmlPluginTemplates.js')

var jsLoaders = ['babel']
var devtool = 'cheap-source-map'

var entryList = [
	'babel-polyfill',
	process.cwd() + '/src/js/main.js'
]

if (process.env.NODE_ENV === 'hot') {
	entryList.unshift('webpack-dev-server/client?http://0.0.0.0:' + config.PORT)
	entryList.unshift('webpack/hot/dev-server')
	jsLoaders.unshift('react-hot')
	devtool = 'eval'
}

module.exports = {
	devtool: devtool,

	entry: {
		'application.bundle': entryList,
		vendor: [
			'babel-polyfill',
			'jquery',
			'lodash',
			'react',
			'react-addons-create-fragment',
			'react-dom',
			'react-router'
		]
	},

	output: {
		path: process.cwd() + '/dest',
		filename: '[name].min.js',
		publicPath: getPublicPath()
	},

	plugins: getPlugins(),

	module: {
		preLoaders: [
			{ test: /\.js$/, loader: 'eslint-loader', exclude: [/node_modules/, /semantic/] }
		],

		loaders: [
			{ test: /\.js$/, loaders: jsLoaders, exclude: /node_modules/  },
			{ test: /\.less$/, loader: 'style/useable!css!less', exclude: /node_modules/ },
			{ test: /\.json$/, loader: 'json' },
			{ test: /\.css$/, exclude: /\.useable\.css$/, loader: 'style!css' },
			{ test: /\.useable\.css$/, loader: 'style/useable!css' },
			{ test: /\.useable\.less$/, loader: 'style/useable!css!less' },
			{
				test: /\.(eot|woff|woff2|ttf|svg)$/,
				loader: 'url-loader?limit=30000&name=[name]-[hash].[ext]'
			},
			{
				test: /\.(jpe?g|png|gif)$/i,
				loaders: [
						'file?hash=sha512&digest=hex&name=[hash].[ext]',
				]
			}
		]
	},

	resolve: {
		root: process.cwd() + '/src',
		extensions: ['', '.js'],
		alias: {
			'': process.cwd() + '/src'
		}
	}
}
function getPlugins() {
 var plugins = [
		new webpack.ProvidePlugin({
			$: 'jquery',
			jQuery: 'jquery',
			'window.jQuery': 'jquery'
		}),
		new webpack.NoErrorsPlugin(),
		new webpack.optimize.CommonsChunkPlugin('vendor', 'vendor.bundle.min.js'),

		new webpack.DefinePlugin({
			'process.env': {
				NODE_ENV: JSON.stringify(process.env.NODE_ENV === 'production' ? 'production' :  'hot')
			}
		}),

		new CleanWebpackPlugin(['dest'], {
			root: process.cwd(),
			verbose: true,
			dry: false
		})
	].concat(
		getHtmlPluginTemplates(
			process.cwd() + '/src/templates',
			process.cwd() + '/dest'
		)
	)

	if (process.env.NODE_ENV === 'hot') {
		plugins.unshift(
			new webpack.HotModuleReplacementPlugin()
		)
	}

	if (process.env.NODE_ENV === 'production') {
		plugins.unshift(
			new webpack.optimize.UglifyJsPlugin({
				compress: {
					warnings: false
				}
			})
		)
	}

	return plugins
}
function getPublicPath() {
	return '/'
}
