import React from 'react'       
import Inside from './Inside'
import { VscDebugStart } from "react-icons/vsc";
import { Textarea } from "@/components/ui/textarea"

import { TfiViewList } from 'react-icons/tfi'

const Body = () => {
  return (
    <div className='w-full max-h-full flex 
     flex-col realtive '><TfiViewList className=' absolute top-16 z-10 md:hidden w-10 h-10' />
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
        <div className='flex bg-slate-700 flex-row items-center  p-4'><Textarea  /><VscDebugStart  className='w-16 h-16 hover:bg-green-400'/></div>
        
    </div>
  )
}

export default Body