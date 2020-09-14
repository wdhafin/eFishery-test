const JWTService = require('../services/auth.service');

// usually: "Authorization: Bearer [token]" or "token: [token]"
module.exports = (req, res, next) => {
  let tokenToVerify;

  if (req.header('Authorization')) {
    const parts = req.header('Authorization').split(' ');

    if (parts.length === 2) {
      const scheme = parts[0];
      const credentials = parts[1];

      if (/^Bearer$/.test(scheme)) {
        tokenToVerify = credentials;
      } else {
        return res.status(401).json({
          message: 'Unauthorized',
          data: null,
        });
      }
    } else {
      return res.status(401).json({
        message: 'Unauthorized',
        data: null,
      });
    }
  } else if (req.body.token) {
    tokenToVerify = req.body.token;
    delete req.query.token;
  } else {
    return res.status(401).json({
      message: 'Unauthorized',
      data: null,
    });
  }

  return JWTService().verify(tokenToVerify, async (err, thisToken) => {
    if (err) {
      return res.status(401).json({
        message: err.message,
        data: null,
      });
    }
    // const user = await User.findByPk(thisToken.id, {
    //   include: [{ model: Role, required: true }],
    //   logging: false,
    // });
    // // console.log("user", user);
    // if (user) {
    //   req.User = user;
    return next();
    // } else {
    //   return res.status(401).json({
    //     message: 'Unauthorized',
    //     data: null,
    //   });
    // }
  });
};
