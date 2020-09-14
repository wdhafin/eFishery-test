/**
 * third party libraries
 */
const bodyParser = require('body-parser');
const express = require('express');
const http = require('http');
const moment = require('moment');
const mapRoutes = require('express-routes-mapper');
require('dotenv').config();

/**
 * server configuration
 */
const config = require('../config/');
const errorService = require('./services/error.service')();
const LoggerServive = require('./services/logger.service')();
const auth = require('./policies/auth.policy');

/**
 * express application
 */
const app = express();
const server = http.Server(app);

// The request handler must be the first middleware on the app
const mappedAuthRoutes = mapRoutes(config.privateRoutes, 'api/controllers/');

// parsing the request bodys
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

// secure your private routes with jwt authentication middleware
app.all('/private/*', (req, res, next) => auth(req, res, next));

// fill routes for express application
app.use('/private', mappedAuthRoutes);

// The error handler must be before any other error middleware and after all controllers
app.use(errorService.errorHandler);

server.listen(config.port, () => {
  LoggerServive.success(
    `Server is running on port: ${config.port} at ${moment()
      .local()
      .format('DD MMMM, YYYY HH:mm:ss Z')}`
  );
  LoggerServive.success(
    `Server is running on Node Version: ${process.version}`
  );
});
