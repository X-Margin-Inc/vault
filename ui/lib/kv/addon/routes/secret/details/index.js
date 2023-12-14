/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Route from '@ember/routing/route';
import { breadcrumbsForSecret } from 'kv/utils/kv-breadcrumbs';

export default class KvSecretDetailsIndexRoute extends Route {
  setupController(controller, resolvedModel) {
    super.setupController(controller, resolvedModel);

    const breadcrumbsArray = [
      { label: 'secrets', route: 'secrets', linkExternal: true },
      { label: resolvedModel.backend, route: 'list' },
      ...breadcrumbsForSecret(resolvedModel.path, true),
    ];
<<<<<<< HEAD
=======

>>>>>>> 4cb759cfc9 (fixed log)
    controller.breadcrumbs = breadcrumbsArray;
  }
}
