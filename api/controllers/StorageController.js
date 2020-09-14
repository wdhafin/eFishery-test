const axios = require('axios');

const StorageController = () => {
  const list = async (req, res, next) => {
    axios
      .get(
        'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list'
      )
      .then((response) => {
        const storages = response.data;
        let currencyRate;
        axios
          .get(
            'https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=a640035639d0c4eb3165'
          )
          .then((response) => {
            currencyRate = response.data.IDR_USD;
            let storagesData = [];
            storages.forEach((storage) => {
              usd = Math.round(storage.price * currencyRate * 100) / 100;
              storage = { ...storage, price_usd: usd };
              storagesData.push(storage);
            });
            res.json({
              data: storagesData,
            });
          })
          .catch((error) => {
            console.log(error);
          });
      })
      .catch((error) => {
        next(error);
      });
  };
  const listAggregate = async (req, res, next) => {
    axios
      .get(
        'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list'
      )
      .then((response) => {
        const storages = response.data;
        res.json({
          data: storages,
        });
      })
      .catch((error) => {
        next(error);
      });
  };

  return {
    list,
    listAggregate,
  };
};

module.exports = StorageController;
