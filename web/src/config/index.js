var ENV = process.env.NODE_ENV || 'hot'

function selectByEnv(map){
	return map[ENV]
}

module.exports = {
	NAME: '<app-name>-' + ENV,

	PORT: selectByEnv({
		hot: 8080,
		production: 8080
	})
}
