import js from '@eslint/js'
import globals from 'globals'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'
import tseslint from 'typescript-eslint'
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'
import i18next from 'eslint-plugin-i18next'
import jsxA11y from 'eslint-plugin-jsx-a11y'

export default tseslint.config(
  { ignores: ['dist', 'dev-dist'] },
  {
    extends: [
      js.configs.recommended,
      ...tseslint.configs.stylisticTypeChecked,
      ...tseslint.configs.recommendedTypeChecked,
      {
        languageOptions: {
          parserOptions: {
            projectService: true,
            tsconfigRootDir: import.meta.dirname,
          },
        },
      },
    ],
    files: ['**/*.{ts,tsx}'],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      'react-x': reactX,
      'react-dom': reactDom,
      'jsx-a11y': jsxA11y,
    },
    rules: {
      ...reactX.configs['recommended-typescript'].rules,
      ...reactDom.configs.recommended.rules,
      ...reactHooks.configs['recommended-latest'].rules,
      ...jsxA11y.flatConfigs.strict.rules,
      // Safari/VoiceOver drops list semantics when list-style is none, so role="list"
      // on a list-none <ul> is a workaround rather than redundancy.
      'jsx-a11y/no-redundant-roles': ['error', { ul: ['list'] }],
      'jsx-a11y/no-noninteractive-element-interactions': [
        'error',
        {
          body: ['onError', 'onLoad'],
          iframe: ['onError', 'onLoad'],
          img: ['onError', 'onLoad'],
          // The form-level onBlur handlers delegate from the natively focusable
          // controls inside; the <form> itself is never a focus target.
          form: ['onBlur', 'onFocus'],
        },
      ],
      'react-refresh/only-export-components': [
        'warn',
        { allowConstantExport: true, allowExportNames: ['Route', 'loader'] },
      ],
      '@typescript-eslint/no-unused-vars': [
        'error',
        { argsIgnorePattern: '^_', varsIgnorePattern: '^_' },
      ],
      '@typescript-eslint/only-throw-error': [
        'error',
        {
          allow: [
            {
              from: 'package',
              name: 'Redirect',
              package: '@tanstack/router',
            },
            {
              from: 'package',
              name: 'NotFoundError',
              package: '@tanstack/router',
            },
          ],
        },
      ],
      '@typescript-eslint/no-empty-object-type': [
        'error',
        { allowInterfaces: 'with-single-extends' },
      ],
      '@typescript-eslint/no-misused-promises': [
        'error',
        {
          checksVoidReturn: false,
        },
      ],
    },
  },
  {
    files: ['src/routes/**/*.tsx'],
    rules: {
      'react-refresh/only-export-components': 'off',
    },
  },
  {
    files: ['src/**/*.{ts,tsx}'],
    plugins: { i18next },
    rules: {
      // `jsx-text-only` catches literal text nodes (`<p>Rohtext</p>`) but NOT string
      // attributes such as `aria-label="…"`, `title="…"` or `placeholder="…"` — which is
      // where most of this plan's extracted copy actually lived. Widening the mode was
      // evaluated and rejected: without an allowlist it produced 113 findings, mostly
      // noise (class names, test ids, non-UI strings). This rule is a floor, not full
      // coverage; grep for `aria-label="` and `placeholder = '` to catch the rest.
      'i18next/no-literal-string': [
        'error',
        {
          mode: 'jsx-text-only',
          'should-validate-template': true,
          message: 'User-facing text belongs in a catalog under src/locales.',
        },
      ],
    },
  },
  {
    // main.tsx renders its bootstrap-failure fallback before I18nextProvider
    // mounts, so there is no instance to translate with (unlike
    // ErrorFallback.tsx, which renders after and uses getI18n()).
    files: [
      'src/routes/_protected/debug/**',
      'src/components/debug/**',
      'src/**/*.test.{ts,tsx}',
      'src/main.tsx',
    ],
    rules: { 'i18next/no-literal-string': 'off' },
  },
  {
    // Test files render fixtures and mocks, not shipped markup — an <a> stub standing
    // in for a router link is not an accessibility defect.
    files: ['src/**/*.test.{ts,tsx}', 'src/test/**'],
    rules: Object.fromEntries(Object.keys(jsxA11y.flatConfigs.strict.rules).map((r) => [r, 'off'])),
  },
)
