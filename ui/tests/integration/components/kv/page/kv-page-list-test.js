/**
 * Copyright (c) HashiCorp, Inc.
<<<<<<< HEAD
 * SPDX-License-Identifier: BUSL-1.1
=======
 * SPDX-License-Identifier: MPL-2.0
>>>>>>> 4cb759cfc9 (fixed log)
 */

import { module, test } from 'qunit';
import { setupRenderingTest } from 'ember-qunit';
import { setupEngine } from 'ember-engines/test-support';
import { setupMirage } from 'ember-cli-mirage/test-support';
import { render, click } from '@ember/test-helpers';
import { hbs } from 'ember-cli-htmlbars';
import { kvMetadataPath } from 'vault/utils/kv-path';
import { allowAllCapabilitiesStub } from 'vault/tests/helpers/stubs';
import { PAGE } from 'vault/tests/helpers/kv/kv-selectors';
<<<<<<< HEAD
import { setRunOptions } from 'ember-a11y-testing/test-support';
=======
>>>>>>> 4cb759cfc9 (fixed log)

const CREATE_RECORDS = (number, store, server) => {
  const mirageList = server.createList('kv-metadatum', number, 'withCustomPath');
  mirageList.forEach((record) => {
    record.data.path = record.path;
    record.id = kvMetadataPath(record.data.backend, record.data.path);
    store.pushPayload('kv/metadata', {
      modelName: 'kv/metadata',
      ...record,
    });
  });
};

const META = {
  currentPage: 1,
  lastPage: 2,
  nextPage: 2,
  prevPage: 1,
  total: 16,
  filteredTotal: 16,
  pageSize: 15,
};

module('Integration | Component | kv | Page::List', function (hooks) {
  setupRenderingTest(hooks);
  setupEngine(hooks, 'kv');
  setupMirage(hooks);

  hooks.beforeEach(async function () {
    this.server.post('/sys/capabilities-self', allowAllCapabilitiesStub());
    this.store = this.owner.lookup('service:store');
<<<<<<< HEAD
    setRunOptions({
      rules: {
        // TODO: ConfirmAction renders modal within list when @isInDropdown
        list: { enabled: false },
      },
    });
  });

  test('it renders Pagination and allows you to delete a kv/metadata record', async function (assert) {
    assert.expect(20);
    CREATE_RECORDS(15, this.store, this.server);
    this.model = await this.store.peekAll('kv/metadata');
    this.model.meta = META;
    this.backend = 'kv-engine';
    this.breadcrumbs = [
      { label: 'secrets', route: 'secrets', linkExternal: true },
      { label: this.backend, route: 'list' },
=======
  });

  test('it renders Pagination and allows you to delete a kv/metadata record', async function (assert) {
    assert.expect(19);
    CREATE_RECORDS(15, this.store, this.server);
    this.model = await this.store.peekAll('kv/metadata');
    this.model.meta = META;
    this.breadcrumbs = [
      { label: 'secrets', route: 'secrets', linkExternal: true },
      { label: this.model.backend, route: 'list' },
>>>>>>> 4cb759cfc9 (fixed log)
    ];
    this.failedDirectoryQuery = false;
    await render(
      hbs`<Page::List
<<<<<<< HEAD
      @secrets={{this.model}}
      @backend={{this.backend}}
      @failedDirectoryQuery={{this.failedDirectoryQuery}}
      @breadcrumbs={{this.breadcrumbs}}
=======
      @secrets={{this.model}} 
      @backend={{this.model.backend}}
      @failedDirectoryQuery={{this.failedDirectoryQuery}}
      @breadcrumbs={{this.breadcrumbs}} 
>>>>>>> 4cb759cfc9 (fixed log)
      @meta={{this.model.meta}}
    />`,
      {
        owner: this.engine,
      }
    );
<<<<<<< HEAD

    assert.dom(PAGE.list.pagination).exists('shows hds pagination component');
    assert.dom(PAGE.list.paginationInfo).hasText('1–15 of 16', 'shows correct page of pages');
    assert.dom(PAGE.title).includesText(this.backend, 'shows backend as title');
=======
    assert.dom(PAGE.list.pagination).exists('shows hds pagination component');
    assert.dom(PAGE.list.paginationInfo).hasText('1–15 of 16', 'shows correct page of pages');
>>>>>>> 4cb759cfc9 (fixed log)

    this.model.forEach((record) => {
      assert.dom(PAGE.list.item(record.path)).exists('lists all records from 0-14 on the first page');
    });

    this.server.delete(kvMetadataPath('kv-engine', 'my-secret-0'), () => {
      assert.ok(true, 'request made to correct endpoint on delete metadata.');
    });

    const popupSelector = `${PAGE.list.item('my-secret-0')} ${PAGE.popup}`;
    await click(popupSelector);
    await click('[data-test-confirm-action-trigger]');
    await click('[data-test-confirm-button]');
    assert.dom(PAGE.list.item('my-secret-0')).doesNotExist('deleted the first record from the list');
  });
});
