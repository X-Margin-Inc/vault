/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Store from '@ember-data/store';
import { AdapterRegistry } from 'ember-data/adapter';

export default interface LdapRoleAdapter extends AdapterRegistry {
  fetchCredentials(backend: string, type: string, name: string);
  rotateStaticPassword(backend: string, name: string);
}
