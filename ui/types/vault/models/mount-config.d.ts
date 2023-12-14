/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Model from '@ember-data/model';

export default class MountConfigModel extends Model {
  defaultLeaseTtl: string;
  maxLeaseTtl: string;
  auditNonHmacRequestKeys: string;
  auditNonHmacResponseKeys: string;
  listingVisibility: string;
  passthroughRequestHeaders: string;
  allowedResponseHeaders: string;
  tokenType: string;
  allowedManagedKeys: string;
}
