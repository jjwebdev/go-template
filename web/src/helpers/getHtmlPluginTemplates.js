var HtmlWebpackPlugin = require('html-webpack-plugin')
var fs = require('fs')

/**
 * Helper method that returns the HtmlWebpackPlugin instantiations to
 * transform ejs templates into html.  Other files types are left as is.
 *
 * @param {string} sourcePath - The source template path.
 * @param {string} destPath - The destination outpute path.
 **/
module.exports = function getHtmlPluginTemplates(sourcePath, destPath) {
	if (!sourcePath || !destPath) return []

	var files = fs.readdirSync(sourcePath),
			templates = []

	files.forEach(function(file) {
		if (file.slice(-3) === 'ejs') {
			templates.push(
				new HtmlWebpackPlugin({
					template: sourcePath + '/' + file,
					filename: destPath + '/' + file.slice(0, -4) + '.html',
					inject: true
				})
			)
		} else {
			templates.push(
				new HtmlWebpackPlugin({
					template: sourcePath + '/' + file,
					filename: destPath + '/' + file,
					inject: false
				})
			)
		}
	})

	return templates
}
