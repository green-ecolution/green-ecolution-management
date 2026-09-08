import { defineConfig } from 'vitest/config'
import type { Plugin } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'
import wasm from 'vite-plugin-wasm'
import topLevelAwait from 'vite-plugin-top-level-await'

// vite-plugin-wasm only inlines the module as a base64 data URI — the one form
// that loads under node — when it finds a plugin literally named "vitest".
// Vitest 5 renamed its plugins to the "vitest:*" namespace, so the detection
// fails and the plugin emits a dev-server /@fs/ URL that fetch() rejects as
// invalid. This marker restores the detection until the plugin ships a fix.
const vitestMarker: Plugin = { name: 'vitest' }

export default defineConfig({
  plugins: [vitestMarker, react(), wasm(), topLevelAwait()],
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: ['./src/test/setup.ts'],
    include: ['src/**/*.{test,spec}.{ts,tsx}', '../handbook/src/**/*.test.mjs'],
    exclude: ['node_modules', 'dist'],
    css: true,
    coverage: {
      provider: 'v8',
      reporter: ['text', 'json', 'html'],
      include: ['src/**/*.{ts,tsx}'],
      exclude: ['src/test/**', 'src/routes/**', '**/*.d.ts', 'src/routeTree.gen.ts'],
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      'virtual:pwa-register/react': path.resolve(__dirname, './src/test/mocks/pwaRegister.ts'),
    },
  },
})
