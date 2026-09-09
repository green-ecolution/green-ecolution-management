import { createContext, useContext, useEffect, useState, type PropsWithChildren } from 'react'
import { connectToHost, type PluginContext } from '../handshake'

const Context = createContext<PluginContext | undefined>(undefined)

/**
 * Renders children only once the host has answered the handshake, so a
 * consumer calling usePluginContext() always gets the real context rather
 * than having to handle a not-yet-connected state itself.
 */
export const PluginProvider = ({ children }: PropsWithChildren) => {
  const [context, setContext] = useState<PluginContext>()

  useEffect(() => {
    let cancelled = false
    void connectToHost().then((received) => {
      if (!cancelled) setContext(received)
    })
    return () => {
      cancelled = true
    }
  }, [])

  if (context === undefined) return null

  return <Context.Provider value={context}>{children}</Context.Provider>
}

export const usePluginContext = () => {
  const context = useContext(Context)
  if (context === undefined) {
    throw new Error('usePluginContext must be used within a PluginProvider')
  }
  return context
}
