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
import { render } from '@ember/test-helpers';
import { hbs } from 'ember-cli-htmlbars';
import { PAGE } from 'vault/tests/helpers/kv/kv-selectors';
/* eslint-disable no-useless-escape */

module('Integration | Component | kv-v2 | Page::Secret::Paths', function (hooks) {
  setupRenderingTest(hooks);
  setupEngine(hooks, 'kv');

  hooks.beforeEach(async function () {
    this.backend = 'kv-engine';
    this.path = 'my-secret';
    this.breadcrumbs = [
      { label: 'secrets', route: 'secrets', linkExternal: true },
      { label: this.backend, route: 'list' },
      { label: this.path },
    ];

    this.assertClipboard = (assert, element, expected) => {
      assert.dom(element).hasAttribute('data-test-copy-button', expected);
    };
  });

  test('it renders copyable paths', async function (assert) {
    assert.expect(6);

    const paths = [
      { label: 'API path', expected: `/v1/${this.backend}/data/${this.path}` },
      { label: 'CLI path', expected: `-mount="${this.backend}" "${this.path}"` },
      { label: 'API path for metadata', expected: `/v1/${this.backend}/metadata/${this.path}` },
    ];

    await render(
      hbs`
<Page::Secret::Paths
  @path={{this.path}}
  @backend={{this.backend}}
  @breadcrumbs={{this.breadcrumbs}}
/>
      `,
      { owner: this.engine }
    );

    for (const path of paths) {
      assert.dom(PAGE.infoRowValue(path.label)).hasText(path.expected);
      this.assertClipboard(assert, PAGE.paths.copyButton(path.label), path.expected);
    }
  });

  test('it renders copyable encoded mount and secret paths', async function (assert) {
    assert.expect(6);
    this.path = `my spacey!"secret`;
    this.backend = `my fancy!"backend`;
    const backend = encodeURIComponent(this.backend);
    const path = encodeURIComponent(this.path);
    const paths = [
      {
        label: 'API path',
        expected: `/v1/${backend}/data/${path}`,
      },
      { label: 'CLI path', expected: `-mount="${this.backend}" "${this.path}"` },
      {
        label: 'API path for metadata',
        expected: `/v1/${backend}/metadata/${path}`,
      },
    ];

    await render(
      hbs`
<Page::Secret::Paths
  @path={{this.path}}
  @backend={{this.backend}}
  @breadcrumbs={{this.breadcrumbs}}
/>
      `,
      { owner: this.engine }
    );

    for (const path of paths) {
      assert.dom(PAGE.infoRowValue(path.label)).hasText(path.expected);
      this.assertClipboard(assert, PAGE.paths.copyButton(path.label), path.expected);
    }
  });

  test('it renders copyable commands', async function (assert) {
<<<<<<< HEAD
    const url = `https://127.0.0.1:8200/v1/${this.backend}/data/${this.path}`;
    const expected = {
      cli: `vault kv get -mount="${this.backend}" "${this.path}"`,
      api: `curl \\ --header \"X-Vault-Token: ...\" \\ --request GET \\ ${url}`,
=======
    assert.expect(4);
    const url = `https://127.0.0.1:8200/v1/${this.backend}/data/${this.path}`;
    const expected = {
      cli: `vault kv get -mount="${this.backend}" "${this.path}"`,
      apiDisplay: `curl \\ --header \"X-Vault-Token: ...\" \\ --request GET \\ ${url}`,
      apiCopy: `curl  --header \"X-Vault-Token: ...\"  --request GET \ ${url}`,
>>>>>>> 4cb759cfc9 (fixed log)
    };
    await render(
      hbs`
<Page::Secret::Paths
  @path={{this.path}}
  @backend={{this.backend}}
  @breadcrumbs={{this.breadcrumbs}}
/>
      `,
      { owner: this.engine }
    );

    assert.dom(PAGE.paths.codeSnippet('cli')).hasText(expected.cli);
<<<<<<< HEAD
    assert.dom(PAGE.paths.codeSnippet('api')).hasText(expected.api);
  });

  test('it renders copyable encoded mount and path commands', async function (assert) {
=======
    assert.dom(PAGE.paths.snippetCopy('cli')).hasAttribute('data-test-copy-button', expected.cli);
    assert.dom(PAGE.paths.codeSnippet('api')).hasText(expected.apiDisplay);
    assert.dom(PAGE.paths.snippetCopy('api')).hasAttribute('data-test-copy-button', expected.apiCopy);
  });

  test('it renders copyable encoded mount and path commands', async function (assert) {
    assert.expect(4);
>>>>>>> 4cb759cfc9 (fixed log)
    this.path = `my spacey!"secret`;
    this.backend = `my fancy!"backend`;

    const backend = encodeURIComponent(this.backend);
    const path = encodeURIComponent(this.path);
    const url = `https://127.0.0.1:8200/v1/${backend}/data/${path}`;

    const expected = {
      cli: `vault kv get -mount="${this.backend}" "${this.path}"`,
<<<<<<< HEAD
      api: `curl \\ --header \"X-Vault-Token: ...\" \\ --request GET \\ ${url}`,
=======
      apiDisplay: `curl \\ --header \"X-Vault-Token: ...\" \\ --request GET \\ ${url}`,
      apiCopy: `curl  --header \"X-Vault-Token: ...\"  --request GET \ ${url}`,
>>>>>>> 4cb759cfc9 (fixed log)
    };
    await render(
      hbs`
<Page::Secret::Paths
  @path={{this.path}}
  @backend={{this.backend}}
  @breadcrumbs={{this.breadcrumbs}}
/>
      `,
      { owner: this.engine }
    );

    assert.dom(PAGE.paths.codeSnippet('cli')).hasText(expected.cli);
<<<<<<< HEAD
    assert.dom(PAGE.paths.codeSnippet('api')).hasText(expected.api);
=======
    assert.dom(PAGE.paths.snippetCopy('cli')).hasAttribute('data-test-copy-button', expected.cli);
    assert.dom(PAGE.paths.codeSnippet('api')).hasText(expected.apiDisplay);
    assert.dom(PAGE.paths.snippetCopy('api')).hasAttribute('data-test-copy-button', expected.apiCopy);
>>>>>>> 4cb759cfc9 (fixed log)
  });
});
