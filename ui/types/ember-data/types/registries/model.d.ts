/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import Model from '@ember-data/model';
import KvSecretDataModel from 'vault/models/kv/data';
import KvSecretMetadataModel from 'vault/models/kv/metadata';
import PkiActionModel from 'vault/models/pki/action';
import PkiCertificateGenerateModel from 'vault/models/pki/certificate/generate';

declare module 'ember-data/types/registries/model' {
  export default interface ModelRegistry {
    'pki/action': PkiActionModel;
    'pki/certificate/generate': PkiCertificateGenerateModel;
<<<<<<< HEAD
    'kv/data': KvSecretDataModel;
=======
    'kv/data': KvSecretDataModelModel;
>>>>>>> 4cb759cfc9 (fixed log)
    'kv/metadata': KvSecretMetadataModel;
    // Catchall for any other models
    [key: string]: any;
  }
}
