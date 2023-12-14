/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Route from '@ember/routing/route';
import { inject as service } from '@ember/service';

import type RouterService from '@ember/routing/router-service';

export default class LdapLibraryRoute extends Route {
  @service declare readonly router: RouterService;

  redirect() {
    this.router.transitionTo('vault.cluster.secrets.backend.ldap.libraries.library.details');
  }
}
