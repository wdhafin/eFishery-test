const storage = {
  'GET /storages': 'StorageController.list',
  'GET /storages/aggregate': 'StorageController.listAggregate',
};

const util = {
  'GET /util/extract_jwt': 'UtilController.extractjwt',
};

const privateRoutes = {
  ...storage,
  ...util,
};

module.exports = privateRoutes;
