import React from 'react'
import { VscAccount } from "react-icons/vsc";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import "../../index.css"

const Navbar = () => {
    return (
        <div className='z-10  top-0 w-full py-3 flex flex-row bg-slate-400'>
            <div className='rounded-full w-8 h-8 bg-blue-600 '></div>

            <nav className='flex flex-row w-full justify-evenly'>
                <div className=' border-b-violet-500 button-transition'>home</div>
                <div className=' border-b-violet-500 button-transition'>sass</div>
                <div className=' border-b-violet-500 button-transition'>students</div>
            </nav>
            <Avatar><AvatarImage src="https://github.com/shadcn.png"/>
                <AvatarFallback><VscAccount /></AvatarFallback>
            </Avatar>
            </div>

   
    )
}

export default Navbar