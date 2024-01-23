import React from 'react'
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { VscAccount } from 'react-icons/vsc'

const Inside = () => {
  return (
<div className='flex flex-row p-2 gap-2 justify-center items-center w-full h-fit bg-green-300'>
<Avatar>
  <AvatarImage src="https://github.com/shadcn.png"/>
                <AvatarFallback><VscAccount /></AvatarFallback>
            </Avatar>
    <div className='flex flex-col w-full h-fit gap-2 p-2 '>
        <li className=''>name</li>
        <li>class</li>
    </div>
</div>
  )
}

export default Inside