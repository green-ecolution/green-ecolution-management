import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/trees')({
  component: Trees,
})

function Trees() {
  return <div>Hello /trees!</div>
}
