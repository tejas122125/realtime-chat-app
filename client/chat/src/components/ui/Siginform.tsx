import { SigninValidation } from '@/lib/validation';
import { zodResolver } from '@hookform/resolvers/zod';
import React from 'react'
import { useForm } from "react-hook-form"
import { z } from 'zod';
import {Form,FormControl, FormDescription,FormField,FormItem,FormLabel,FormMessage,
  } from "@/components/ui/form"
import { Input } from './input';
import { Button } from './button';
const Siginform = () => {

    const form = useForm<z.infer<typeof SigninValidation>>({
        resolver: zodResolver(SigninValidation),
        defaultValues: {
          email: "",
          password: "",
        },
      });

      function onSubmit(values: z.infer<typeof SigninValidation>) {
       
        console.log(values)
      }

  return (

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}  className="flex flex-col gap-5 w-2/4 mt-4">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className='text-white'>name</FormLabel>
                  <FormControl>
                    <Input autoComplete='off' placeholder="name" {...field} />
                  </FormControl>
                  <FormDescription className='text-white'>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className='text-white'>email</FormLabel>
                  <FormControl>
                    <Input  placeholder="email" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className='text-white'>password</FormLabel>
                  <FormControl>
                    <Input autoComplete='off' placeholder="password" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit">Submit</Button>
          </form>
        </Form>
  )
}

export default Siginform