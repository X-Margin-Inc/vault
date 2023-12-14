/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import ApplicationSerializer from '../application';

export default class LdapRoleSerializer extends ApplicationSerializer {
  primaryKey = 'name';

  serialize(snapshot) {
    // remove all fields that are not relevant to specified role type
    const { fieldsForType } = snapshot.record;
    const json = super.serialize(...arguments);
    Object.keys(json).forEach((key) => {
      if (!fieldsForType.includes(key)) {
        delete json[key];
      }
    });
    return json;
  }
}
