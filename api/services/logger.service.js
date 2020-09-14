const chalk = require('chalk');

const loggerService = () => {
  const error = (...args) => {
    console.error(chalk.bold.red(...args));
  };

  const success = (...args) => {
    console.info(chalk.bold.green(...args));
  };

  const info = (...args) => {
    console.info(chalk.bold.blue(...args));
  };

  return {
    error,
    success,
    info,
  };
};

module.exports = loggerService;
