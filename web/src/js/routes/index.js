const NODE_ENV = process.env.NODE_ENV || 'hot'

if (NODE_ENV === 'hot') {
	module.exports = require('./hot.js')
} else {
	module.exports = require('./production.js')
}
