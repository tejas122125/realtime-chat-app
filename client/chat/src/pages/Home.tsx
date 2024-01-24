import Body from '@/components/ui/Body'
import Navbar from '@/components/ui/Navbar'
import Sidebar from '@/components/ui/Sidebar'
import React from 'react'

const Home = () => {
  return (
<>
    <Navbar />
    <div className='max-h-full relative pt-16 gap-4 flex flex-row'>
      <Sidebar />
      <Body />
    </div>
</>
  )
}

export default Home