import React from 'react'
import { TfiViewList } from "react-icons/tfi";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { VscAccount } from 'react-icons/vsc'
const itemlist = ["m","t","r","tgedfthgfd"]
const Sidebar = () => {
  return (
    <div className='hidden md:flex flex-col min-w-[200px] max-w-fit h-full
     bg-green-200'>sidebaxxxxxxxxxx
    <ul className='flex-col gap-4'>{
       itemlist.map((index,value)=>{
        return (<li className='w-full flex  items-center flex-row gap-4 mt-4 mb-4  rounded-lg justify-start p-4 bg-purple-400'>
            <Avatar><AvatarImage src="https://github.com/shadcn.png"/>
                <AvatarFallback><VscAccount /></AvatarFallback>
            </Avatar>{value}</li>)
       })
    }</ul></div>
  )
}

export default Sidebar