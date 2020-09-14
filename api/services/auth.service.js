const jwt = require('jsonwebtoken');

const secret = 'eFisheryTest';

const authService = () => {
  const issue = (payload) => {
    return jwt.sign(payload, secret, { expiresIn: 10800 * 240 });
  };
  const verify = (token, cb) => jwt.verify(token, secret, {}, cb);

  return {
    issue,
    verify,
  };
};

module.exports = authService;
