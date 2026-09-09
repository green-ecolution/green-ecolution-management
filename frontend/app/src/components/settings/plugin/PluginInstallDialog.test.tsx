import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, expect, it, vi } from 'vitest'
import { PluginInstallDialog } from './PluginInstallDialog'

describe('PluginInstallDialog', () => {
  it('mirrors the plugin permissions into the access permissions until they are edited', async () => {
    const user = userEvent.setup()
    render(<PluginInstallDialog open onSubmit={vi.fn()} onOpenChange={vi.fn()} />)

    await user.click(screen.getByRole('checkbox', { name: /bäume.*anlegen/i }))

    expect(screen.getByRole('checkbox', { name: /zugang.*bäume.*anlegen/i })).toBeChecked()
  })

  it('stops mirroring once the access permissions were touched', async () => {
    const user = userEvent.setup()
    render(<PluginInstallDialog open onSubmit={vi.fn()} onOpenChange={vi.fn()} />)

    await user.click(screen.getByRole('checkbox', { name: /zugang.*bäume.*lesen/i }))
    await user.click(screen.getByRole('checkbox', { name: /bäume.*anlegen/i }))

    expect(screen.getByRole('checkbox', { name: /zugang.*bäume.*anlegen/i })).not.toBeChecked()
  })

  it('rejects a non-https external url', async () => {
    const user = userEvent.setup()
    const onSubmit = vi.fn()
    render(<PluginInstallDialog open onSubmit={onSubmit} onOpenChange={vi.fn()} />)

    await user.type(screen.getByLabelText(/slug/i), 'acme')
    await user.type(screen.getByLabelText(/name/i), 'Acme')
    await user.selectOptions(screen.getByLabelText(/ansicht/i), 'external')
    await user.type(screen.getByLabelText(/adresse/i), 'http://plugin.example.com')
    await user.click(screen.getByRole('button', { name: /installieren/i }))

    expect(onSubmit).not.toHaveBeenCalled()
    expect(screen.getByText(/https/i)).toBeVisible()
  })
})
