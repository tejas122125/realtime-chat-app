import { z } from "zod";


export const SigninValidation = z.object({
    name : z.string().min(2,{message :"Too Short"}),
    email : z.string().min(2,{message : "Too Short"}),
    password : z.string().min(8,{message : "Too Weak Password"})
  })