/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Engine from '@ember/engine';

import loadInitializers from 'ember-load-initializers';
import Resolver from 'ember-resolver';

import config from './config/environment';

const { modulePrefix } = config;

export default class KvEngine extends Engine {
  modulePrefix = modulePrefix;
  Resolver = Resolver;
  dependencies = {
    services: [
      'download',
      'namespace',
      'router',
      'store',
      'secret-mount-path',
      'flash-messages',
      'control-group',
    ],
<<<<<<< HEAD
    externalRoutes: ['secrets', 'syncDestination'],
=======
    externalRoutes: ['secrets'],
>>>>>>> 4cb759cfc9 (fixed log)
  };
}

loadInitializers(KvEngine, modulePrefix);
