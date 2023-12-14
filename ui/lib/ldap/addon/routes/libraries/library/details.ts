/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Route from '@ember/routing/route';

import type LdapLibraryModel from 'vault/models/ldap/library';
import type Controller from '@ember/controller';
import type Transition from '@ember/routing/transition';
import type { Breadcrumb } from 'vault/vault/app-types';

interface LdapLibraryDetailsController extends Controller {
  breadcrumbs: Array<Breadcrumb>;
  model: LdapLibraryModel;
}

export default class LdapLibraryDetailsRoute extends Route {
  setupController(
    controller: LdapLibraryDetailsController,
    resolvedModel: LdapLibraryModel,
    transition: Transition
  ) {
    super.setupController(controller, resolvedModel, transition);

    controller.breadcrumbs = [
      { label: resolvedModel.backend, route: 'overview' },
      { label: 'libraries', route: 'libraries' },
      { label: resolvedModel.name },
    ];
  }
}
