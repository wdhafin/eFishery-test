const UtilController = () => {
  const extractjwt = async (req, res, next) => {
    token = req.Token;
    res.json({
      name: token.name,
      phone: token.phone,
      role: token.role,
      timestamp: token.timestamp,
    });
  };
  return {
    extractjwt,
  };
};

module.exports = UtilController;
