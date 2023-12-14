/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import Component from '@glimmer/component';
import { inject as service } from '@ember/service';

export default class SidebarNavClusterComponent extends Component {
  @service currentCluster;
  @service version;
  @service auth;
  @service namespace;

  get cluster() {
    return this.currentCluster.cluster;
  }
<<<<<<< HEAD

  get isRootNamespace() {
    // should only return true if we're in the true root namespace
    return this.namespace.inRootNamespace && !this.currentCluster.hasChrootNamespace;
  }
=======
>>>>>>> 4cb759cfc9 (fixed log)
}
