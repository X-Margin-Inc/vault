/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import { module, test } from 'qunit';
import { setupRenderingTest } from 'ember-qunit';
import { render } from '@ember/test-helpers';
import hbs from 'htmlbars-inline-precompile';

const CARD_TITLE = 'Card title';
const ACTION_TEXT = 'View card';
const SUBTEXT = 'This is subtext for card';

<<<<<<< HEAD
const SELECTORS = {
  title: '[data-test-overview-card-title]',
  subtitle: '[data-test-overview-card-subtitle]',
  action: '[data-test-action-text]',
};

=======
>>>>>>> 4cb759cfc9 (fixed log)
module('Integration | Component overview-card', function (hooks) {
  setupRenderingTest(hooks);

  hooks.beforeEach(function () {
    this.set('cardTitle', CARD_TITLE);
    this.set('actionText', ACTION_TEXT);
    this.set('subText', SUBTEXT);
  });

  test('it returns card title, ', async function (assert) {
    await render(hbs`<OverviewCard @cardTitle={{this.cardTitle}}/>`);
<<<<<<< HEAD
    assert.dom(SELECTORS.title).hasText('Card title');
  });
  test('it returns card subtext, ', async function (assert) {
    await render(hbs`<OverviewCard @cardTitle={{this.cardTitle}}  @subText={{this.subText}} />`);
    assert.dom(SELECTORS.subtitle).hasText('This is subtext for card');
=======
    const titleText = this.element.querySelector('h3').innerText;
    assert.strictEqual(titleText, 'Card title');
  });
  test('it returns card subtext, ', async function (assert) {
    await render(hbs`<OverviewCard @cardTitle={{this.cardTitle}}  @subText={{this.subText}} />`);
    const titleText = this.element.querySelector('p').innerText;
    assert.strictEqual(titleText, 'This is subtext for card');
>>>>>>> 4cb759cfc9 (fixed log)
  });
  test('it returns card action text', async function (assert) {
    await render(
      hbs`<OverviewCard @cardTitle={{this.cardTitle}} @actionText={{this.actionText}} @actionTo="route"/>`
    );
<<<<<<< HEAD
    assert.dom(SELECTORS.action).hasText('View card');
=======
    const titleText = this.element.querySelector('a').innerText;
    assert.strictEqual(titleText, 'View card');
>>>>>>> 4cb759cfc9 (fixed log)
  });
});
