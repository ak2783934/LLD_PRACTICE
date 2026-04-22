3. How do you safely iterate over a map while multiple goroutines are writing to it?


And: 

mutex.lock();
defer mutex.unlock();
for x,y:=range map{
    doOperation(x,y)
}


But this blocks the access to this block of code by any other thread till all the values of the map is read. 

But based on the requirements, we might have to lock the map at the value level or the whole map level. 