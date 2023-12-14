/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import buildRoutes from 'ember-engines/routes';

export default buildRoutes(function () {
  this.route('overview');
  this.route('roles', function () {
    this.route('create');
    this.route('role', { path: '/:type/:name' }, function () {
      this.route('details');
      this.route('edit');
      this.route('credentials');
    });
  });
  this.route('libraries', function () {
    this.route('create');
    this.route('library', { path: '/:name' }, function () {
      this.route('details', function () {
        this.route('accounts');
        this.route('configuration');
      });
      this.route('edit');
      this.route('check-out');
    });
  });
  this.route('configure');
  this.route('configuration');
});
