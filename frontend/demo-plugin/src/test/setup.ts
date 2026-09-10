import '@testing-library/jest-dom/vitest'

// jsdom has no ResizeObserver, and the view installs one to report its height
// to the host.
class ResizeObserverStub {
  observe() {}
  unobserve() {}
  disconnect() {}
}

globalThis.ResizeObserver = ResizeObserverStub as unknown as typeof ResizeObserver
