import './index.css'
import { Route, Routes } from "react-router-dom";
import Login from './pages/Login'
import Home from './pages/Home'
const App =()=> {

  return (
    <main className='w-screen relative h-screen  overflow-hidden flex flex-col '>
      <Routes>
        <Route path="/signin" element={<Login />} />
        {/* <Route path="/" element={<Home />} /> */}
        <Route path='/' element={<Home/>}/>

      </Routes>
    </main>
  )
}

export default App
