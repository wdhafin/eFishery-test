const _ = require('lodash');

const errorService = () => {
  const errorHandler = (err, req, res, next) => {
    console.log('Error: ', err);
    if (err && err.status) {
      res.status(err.status).json(
        Object.assign(
          {
            message: err.message,
            data: err.data,
          },
          err.meta
        )
      );
    } else {
      res.status(500).json({
        message: err.message || 'something_broke',
        data: err,
      });
    }
  };

  return {
    errorHandler,
  };
};

module.exports = errorService;
