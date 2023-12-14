/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Engine from 'ember-engines/engine';
import loadInitializers from 'ember-load-initializers';
import Resolver from 'ember-resolver';
import config from './config/environment';

const { modulePrefix } = config;

export default class LdapEngine extends Engine {
  modulePrefix = modulePrefix;
  Resolver = Resolver;
  dependencies = {
    services: ['router', 'store', 'secret-mount-path', 'flash-messages', 'auth'],
    externalRoutes: ['secrets'],
  };
}

loadInitializers(LdapEngine, modulePrefix);
