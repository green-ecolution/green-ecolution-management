import { createFileRoute } from '@tanstack/react-router'
import MapHeader from '../components/MapHeader'
import Map from '../components/Map'

export const Route = createFileRoute('/')({
  component: Index,
})

function Index() {
  return (
    <div className='relative'>
      <MapHeader />
      <Map />
    </div>
  )
}
