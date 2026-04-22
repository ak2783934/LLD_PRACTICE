4. Describe what defer, panic, and recover are and how you’d use them in production HTTP handlers.


defer: Its used to execute the function called after defer at the end of the function execution. 
basically when the scope of a function about to end, the function written after defer is called. 
It helps us to make sure that certain function which are to be called at the end(like closing a channel, closing the context) are done properly, rather than handeling them manually. 

ex: 

ctx, cancel := context.WithTimeout(context.Background());
defer cancel()

res, err:= doOperation()
if err != nil{
    return;
}
return;


if we don't use defer, then at each return, we might have to write cancel manually, but with defer, its taken care already. 


Panic: 
Panic is signal in golang, that happens when we try to access an address that doesn't exist. After panic, go stop and kill the thread and start a new execution after recoverry. 


Recover: 
