const storage = {
  'GET /storages': 'StorageController.list',
  'GET /storages/aggregate': 'StorageController.listAggregate',
};

const privateRoutes = {
  ...storage,
};

module.exports = privateRoutes;
