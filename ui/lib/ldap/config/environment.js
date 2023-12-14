/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

/* eslint-env node */
'use strict';

module.exports = function (environment) {
  const ENV = {
    modulePrefix: 'ldap',
    environment,
  };

  return ENV;
};
