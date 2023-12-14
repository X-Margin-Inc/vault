/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Route from '@ember/routing/route';

import type LdapRoleModel from 'vault/models/ldap/role';
import type Controller from '@ember/controller';
import type Transition from '@ember/routing/transition';
import type { Breadcrumb } from 'vault/vault/app-types';

interface LdapRoleDetailsController extends Controller {
  breadcrumbs: Array<Breadcrumb>;
  model: LdapRoleModel;
}

export default class LdapRoleEditRoute extends Route {
  setupController(
    controller: LdapRoleDetailsController,
    resolvedModel: LdapRoleModel,
    transition: Transition
  ) {
    super.setupController(controller, resolvedModel, transition);

    controller.breadcrumbs = [
      { label: resolvedModel.backend, route: 'overview' },
      { label: 'roles', route: 'roles' },
      { label: resolvedModel.name },
    ];
  }
}
