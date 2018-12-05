'use strict'

var express = require('express')
var path = require('path')

var app = express()
var bodyParser = require('body-parser')

app.use(express.static(path.join(__dirname, '/public')))

app.use(bodyParser.json())

app.use(bodyParser.urlencoded({extended: true}))

var port = 9000
app.listen(port) // Listen on port defined in config file

//app.use('/api', routes)

app.use(function (req, res, next) {
    // Website you wish to allow to connect
  res.setHeader('Access-Control-Allow-Origin', 'http://localhost:' + port)

    // Request methods you wish to allow
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE')

    // Request headers you wish to allow
  res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type')

    // Pass to next layer of middleware
  next()
})
// Server index.html page when request to the root is made
app.get('/', function (req, res, next) {
  res.sendfile('./public/index.html')
})