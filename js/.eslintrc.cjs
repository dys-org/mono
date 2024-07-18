/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution');

module.exports = {
  root: true,
  extends: [
    'plugin:vue/vue3-recommended',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier/skip-formatting',
  ],
  rules: {
    'vue/no-undef-components': ['error', { ignorePatterns: [] }],
    'vue/v-on-event-hyphenation': 'off',
    'vue/multi-word-component-names': 'off',
    // 'vue/require-default-prop': 'off',
  },
  overrides: [
    {
      files: ['**/e2e/**/*.{test,spec}.{js,ts,jsx,tsx}'],
      extends: ['plugin:playwright/recommended'],
    },
  ],
  parserOptions: {
    ecmaVersion: 'latest',
  },
};
