
import Inside from './components/ui/Inside'
import Navbar from './components/ui/Navbar'
import { Button } from './components/ui/button'
import Sidebar from './components/ui/sidebar'
import './index.css'

function App() {

  return (
    <><div className='w-full min-h-fit h-screen flex  flex-col'>
          <Navbar/>
          <div className='h-screen flex flex-row'>
          <Sidebar/>
    <Inside/>
          </div>
    
    </div>

       
    </>
  )
}

export default App
