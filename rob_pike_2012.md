#Concurrency Is Not Parallelism
## Rob Pike 2012

Go provides:
- concurrent execution (goroutines)
- synchronization and messaging (channels)
- multi-way concurrent control (select)

Concurrency and parallelism is not the same thing

Concurrency: Programming as the composition of independently executing processes

    - Not Linux process --> more general/abstract

Parallelism: Programming as the simultaneous execution of possibly related processes

Conc -> dealing with a lot of things at once | structure
    -> i.e. mous keyboard display disk drives

Par -> doing lots of things at once | execution
    -> vector dot product


Concurrency is a way to structure a program by breaking it into pieces that can be
executed independently.

--- Gopher Problem ---
Gopher to incinerate a pile of books.

1 gopher can only transports so much

add gophers and add tools for the gophers

these gophers need to be given orders to not run into each other

This is parallel. The concurrent composition of process is the adding of gophers and
the tools.

...

Observation:

Improved performance by adding a concurrent procedure to the existing design.

Parallism comes from better concurrent expression of the problem.

Concurrent design is the method of decomposing problems into individual processes

...

Now using go:
goroutines:
    a little bit like threads but... different
    
    go myFunc1() // begins running
    myFunc2() // does not wait for myFunc1 to return to begin

    cheap compared to threads. system takes care of system calls that would have 
    been necessary otherwise

channels:
    typed value that allow goroutines to synchronize and exchange info

    timerChan := make(chan time.Time)
    
    // This is an anonymous function, the last () must be used to call the func

    go func() {
        time.Sleep(deltaT)
        timerChan <- time.Now() // send time on timerChan
    }() 
    // Do something else
    completedAt := <-timerChan

select:
    select statement is like a switch statement, but the decision is based on abilitiy to 
    comunicate rather than equal values

    select {
    case v := <-ch1:
        fmt.Println("channel 1 sends", v)
    case v := <-ch2:
        fmt.Println("channel2 sends", v)
    default : // optional
        fmt.Println("neither channel was ready")
    }

...

goroutines aren't free, but they're very cheap

...

load balancer is implicitly parallel and scalable

No locking, no mutexes...

load balancing problem...

 
 
