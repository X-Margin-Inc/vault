/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Component from '@glimmer/component';
import { inject as service } from '@ember/service';

/**
 * @module DashboardVaultVersionTitle
 * DashboardVaultVersionTitle component are use to display the vault version title and the badges
 *
 * @example
 * ```js
 * <Dashboard::VaultVersionTitle @version={{this.versionSvc}} />
 * ```
 */

export default class DashboardVaultVersionTitle extends Component {
  @service namespace;
}
