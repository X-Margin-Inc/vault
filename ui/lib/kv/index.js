/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

/* eslint-env node */
/* eslint-disable n/no-extraneous-require */
'use strict';

const { buildEngine } = require('ember-engines/lib/engine-addon');

module.exports = buildEngine({
  name: 'kv',

  lazyLoading: {
    enabled: false,
  },

  isDevelopingAddon() {
    return true;
  },
});
