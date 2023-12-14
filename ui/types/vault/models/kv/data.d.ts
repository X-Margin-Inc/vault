/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Model from '@ember-data/model';
import CapabilitiesModel from '../capabilities';

export default class KvSecretDataModel extends Model {
  backend: string;
  path: string;
  secretData: object;
  createdTime: string;
  customMetadata: object;
  deletionTime: string;
  destroyed: boolean;
  versions: object;
  failReadErrorCode: number;
  casVersion: number;
  // apiPaths for capabilities
  dataPath: Promise<CapabilitiesModel>;
  metadataPath: Promise<CapabilitiesModel>;
  deletePath: Promise<CapabilitiesModel>;
  destroyPath: Promise<CapabilitiesModel>;
  undeletePath: Promise<CapabilitiesModel>;

  // Capabilities
  get canDeleteLatestVersion(): boolean;
  get canDeleteVersion(): boolean;
  get canUndelete(): boolean;
  get canDestroyVersion(): boolean;
  get canEditData(): boolean;
  get canReadData(): boolean;
  get canReadMetadata(): boolean;
  get canUpdateMetadata(): boolean;
  get canListMetadata(): boolean;
}
