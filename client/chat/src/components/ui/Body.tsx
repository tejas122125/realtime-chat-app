import React from 'react'       
import Inside from './Inside'

const Body = () => {
  return (
    <div className='w-full max-h-full flex 
     flex-col bg-black'>
        <Inside/>
        <div className='bg-slate-700 h-full
         overflow-y-scroll scroll-smooth
         flex flex-col flex-1'>
            <div className='sender self-end bg-green-400  p-4 m-4 rounded-br-none rounded-r-md w-fit h-fit right-2 border-2 '>i am the sender</div>
            <div className='reciever self-start p-4 m-4 bg-purple-500 rounded-br-none rounded-r-md w-fit h-fit left-2 border-2 '>i am the reciever</div>
            <div className='sender self-end bg-green-400  p-4 m-4 rounded-br-none rounded-r-md w-fit h-fit right-2 border-2 '>i am the sender</div>
            <div className='reciever self-start p-4 m-4 bg-purple-500 rounded-br-none rounded-r-md w-fit h-fit left-2 border-2 '>i am the reciever</div>
            <div className='sender self-end bg-green-400  p-4 m-4 rounded-br-none rounded-r-md w-fit h-fit right-2 border-2 '>i am the sender</div>
            <div className='reciever self-start p-4 m-4 bg-purple-500 rounded-br-none rounded-r-md w-fit h-fit left-2 border-2 '>i am the reciever</div>
            <div className='sender self-end bg-green-400  p-4 m-4 rounded-br-none rounded-r-md w-fit h-fit right-2 border-2 '>i am the sender</div>
            <div className='reciever self-start p-4 m-4 bg-purple-500 rounded-br-none rounded-r-md w-fit h-fit left-2 border-2 '>i am the reciever</div>
            
            
        </div>
        <div className='bg-red-500 m-4 p-4'>hdvbfhvbfhgf</div>
        
    </div>
  )
}

export default Body