var ENV = process.env.NODE_ENV || 'hot'

/**
 * Helper method that returns values based off mapped NODE_ENV values.
 *
 * @param {string} map - a mapping the NODE_ENV values.
 * @return {string} the mapped value.
 **/
module.exports = function selectByEnv(map) {
	if (!map) return
	return map[ENV]
}
