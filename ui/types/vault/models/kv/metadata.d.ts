/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import Model from '@ember-data/model';

<<<<<<< HEAD
export default class KvSecretMetadataModel extends Model {
=======
export default class KvSecretDataModel extends Model {
>>>>>>> 4cb759cfc9 (fixed log)
  backend: string;
  path: string;
  fullSecretPath: string;
  maxVersions: number;
  casRequired: boolean;
  deleteVersionAfter: string;
  customMetadata: object;
  createdTime: string;
  currentVersion: number;
  oldestVersion: number;
  updatedTime: string;
  versions: object;
  // apiPaths for capabilities
  dataPath: Promise<CapabilitiesModel>;
  metadataPath: Promise<CapabilitiesModel>;

<<<<<<< HEAD
  get pathIsDirectory(): boolean;
  get isSecretDeleted(): boolean;
  get sortedVersions(): number[];
  get currentSecret(): { state: string; isDeactivated: boolean };

=======
>>>>>>> 4cb759cfc9 (fixed log)
  // Capabilities
  get canDeleteMetadata(): boolean;
  get canReadMetadata(): boolean;
  get canUpdateMetadata(): boolean;
  get canCreateVersionData(): boolean;
}
