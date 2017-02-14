var http = require('http')
var express = require('express')
var app = express()
var config = require(process.cwd() + '/src/config')

process.title = 'KBS'

app.use(express.static(process.cwd() + '/dest'))

app.get('*', function (req, res) {
	res.sendFile(process.cwd() + '/dest/index.html')
})

var server = app.listen(config.PORT, function() {
	console.log('listening on ' + config.PORT)
})

process.on('uncaughtException', shutDown)
process.on('exit', shutDown)
process.on('SIGINT', shutDown)
process.on('SIGTERM', shutDown)
process.on('SIGQUIT', shutDown)
process.on('SIGUSR2', shutDown)

function shutDown(stackTrace) {
	console.log('stackTrace: ', stackTrace)
	server.close()
	process.exit(1)
}
