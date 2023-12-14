/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import SecretsEnginePathAdapter from 'vault/adapters/secrets-engine-path';

export default class LdapConfigAdapter extends SecretsEnginePathAdapter {
  path = 'config';

  async rotateRoot(backend) {
    return this.ajax(this._getURL(backend, 'rotate-root'), 'POST');
  }
}
