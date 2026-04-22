2. Explain the difference between goroutines, channels, and sync primitives. When would you prefer channels over Mutex?


Ans: First thing, they all have different purpose and usecases. 


Goroutine: Go routine is a light weight thread in golang, which helps us to make things async in a simple go file. 
Basically when we execute a goroutine, the function executiokn doesn't wait for its result and proceed forward. 
Using some wait group and other things, we make sure that the function execution doesn't end till the goroutine is not done with its work, or else what might happen is that function execution is done and goroutine is killed because the scope of the function is killed. 

In goroutine, we need to make sure that they are closed in some time, if they keep on running, we might have goroutine which are always active and may create memory leak. 
Since these are very light weight threads, in golang we can spun upto million goroutines without any effort. 

The scheduling of goroutine is taken care by the goruntime itself. 

Channels: Channels is data type in golang, what helps us to pass data in between multiple goroutine. Just assume that we have multiple goroutines running, but they are dependent on each other for data and action. 

so one goroutine can do its operation and then pass the data into the channel, no that channel is will listened by another goroutine and that goroutine will do the futher action and return the data. 

It also help us acheive concurrency and synchronisation between multiple goroutines. 

Sync primitive: These are certain type of data types, which are considered thread safe in golang. Basically it helps us write code in thread safe manner. 

one of the example of these datatypes is sync.Once(), it helps us with singelton pattern, making sure only one thread is able to call the things defined inside sync.once() and no concurrent threads can call the function together. 

Other examples are sync.Mutex(), which help us to create locks in golang. 


channels over mutex:
Using channels and mutex, we try and acheive concurrency and synchronisation in our system. 

I would preffer channels over mutex, when I have somie work where I have to devide my task into multiple smaller goroutines(fan out and fan in models)
