import { expect } from 'chai'
import getHtmlPluginTemplates from 'helpers/getHtmlPluginTemplates'
var HtmlWebpackPlugin = require('html-webpack-plugin')

describe('getHtmlPluginTemplates', () => {
	it('returns [] when no parameters are passed', () => {
		expect(getHtmlPluginTemplates()).to.be.deep.equal([])
	})

	it('returns [] when one parameter is passed', () => {
		expect(getHtmlPluginTemplates(process.cwd() + '/src/templates')).to.be.deep.equal([])
	})

	it('returns [] when one parameter is passed', () => {
		expect(getHtmlPluginTemplates(undefined, process.cwd() + '/dest')).to.be.deep.equal([])
	})

	it('returns HtmlDefinePlugin instances for templates correctly', () => {
		const templates = getHtmlPluginTemplates(
			`${process.cwd()}/src/templates`,
			`${process.cwd()}/dest`
		)

		const expectedOutput = [
			new HtmlWebpackPlugin({
				template: `${process.cwd()}/src/templates/index.ejs`,
				filename: `${process.cwd()}/dest/index.html`,
				inject: true
			})
		]

		expect(templates).to.deep.equal(expectedOutput)
	})
})
