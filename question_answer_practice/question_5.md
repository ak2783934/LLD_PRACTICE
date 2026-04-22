5. How does Go’s memory management work (stack vs heap, GC basics)? When might you care about escape analysis?

Ans: 
Stack memory is execution based memory, it only stores the data for the variables in the current context. 

Heap memory is the global memory in golang, once declared, it not automatically cleared when function execution is done. 

To clear the heap memory, GC runs and clears the nodes which are not accessible from the root. 

Escape analysis: 

Sometimes when a variable is declared inside a function in golang, intially it takes memory in stack, but if its reference is being returned from that function, after return, the memory of that function shifts to heap, because that data needs to exist after the execution is done. 