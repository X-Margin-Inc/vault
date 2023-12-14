/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import ApplicationAdapter from '../application';
export default class KvConfigAdapter extends ApplicationAdapter {
  namespace = 'v1';

  urlForFindRecord(id) {
    return `${this.buildURL()}/${id}/config`;
  }
}
