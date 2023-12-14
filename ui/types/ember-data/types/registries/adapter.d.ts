/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import Application from 'vault/adapters/application';
import Adapter from 'ember-data/adapter';
import ModelRegistry from 'ember-data/types/registries/model';
import PkiIssuerAdapter from 'vault/adapters/pki/issuer';
import PkiTidyAdapter from 'vault/adapters/pki/tidy';
import LdapRoleAdapter from 'vault/adapters/ldap/role';
import LdapLibraryAdapter from 'vault/adapters/ldap/library';
import KvDataAdapter from 'vault/adapters/kv/data';
import KvMetadataAdapter from 'vault/adapters/kv/metadata';
<<<<<<< HEAD
import SyncDestinationAdapter from 'vault/adapters/sync/destination';
import SyncAssociationAdapter from 'vault/adapters/sync/association';
=======
>>>>>>> 4cb759cfc9 (fixed log)

/**
 * Catch-all for ember-data.
 */
export default interface AdapterRegistry {
  'ldap/library': LdapLibraryAdapter;
  'ldap/role': LdapRoleAdapter;
  'pki/issuer': PkiIssuerAdapter;
  'pki/tidy': PkiTidyAdapter;
  'kv/data': KvDataAdapterAdapter;
  'kv/metadata': KvMetadataAdapter;
<<<<<<< HEAD
  'sync/destination': SyncDestinationAdapter;
  'sync/association': SyncAssociationAdapter;
=======
>>>>>>> 4cb759cfc9 (fixed log)
  application: Application;
  [key: keyof ModelRegistry]: Adapter;
}

export default interface AdapterError extends Error {
  httpStatus: number;
}
