/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Route from '@ember/routing/route';
import { hash } from 'rsvp';

import type LdapLibraryModel from 'vault/models/ldap/library';

export default class LdapLibraryRoute extends Route {
  model() {
    const model = this.modelFor('libraries.library') as LdapLibraryModel;
    return hash({
      library: model,
      statuses: model.fetchStatus(),
    });
  }
}
