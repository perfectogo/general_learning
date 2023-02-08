# RAM (Random Access Memory)
* RAM (random access memory) is a computer's short-term memory, where the data that the processor is currently using is stored.

## What does RAM stand for?
* RAM stands for random access memory, and it’s one of the most fundamental elements of computing. RAM is a temporary memory bank where your computer stores data it needs to retrieve quickly. RAM keeps data easily accessible so your processor can quickly find it without having to go into long-term storage to complete immediate processing tasks.

* Nearly all computers have a way of storing information for longer-term access, too. But the memory needed to run the process you’re currently working on is stored and accessed in your computer’s RAM.

## What does RAM do?
* RAM is a form of temporary storage that gets wiped when you turn your computer off. 

* RAM offers lightning-fast data access, which makes it ideal for the processes, apps, and programs your computer is actively working on, such as the data needed to surf the internet through your web browser.

## RAM is of two types −

1. Static RAM (SRAM)
2. Dynamic RAM (DRAM)

* Static RAM (SRAM)
The word static indicates that the memory retains its contents as long as power is being supplied. However, data is lost when the power gets down due to volatile nature. SRAM chips use a matrix of 6-transistors and no capacitors. Transistors do not require power to prevent leakage, so SRAM need not be refreshed on a regular basis. <br/><br/>
There is extra space in the matrix, hence SRAM uses more chips than DRAM for the same amount of storage space, making the manufacturing costs higher. SRAM is thus used as cache memory and has very fast access.

    * Characteristic of Static RAM
        1. Long life
        1. No need to refresh
        1. Faster
        1. Used as cache memory
        1. Large size
        1. Expensive
        1. High power consumption

* Dynamic RAM (DRAM)
DRAM, unlike SRAM, must be continually refreshed in order to maintain the data. This is done by placing the memory on a refresh circuit that rewrites the data several hundred times per second. DRAM is used for most system memory as it is cheap and small. All DRAMs are made up of memory cells, which are composed of one capacitor and one transistor. 
    * haracteristics of Dynamic RAM
        1. Short data lifetime
        1. Needs to be refreshed continuously
        1. Slower as compared to SRAM
        1. Used as RAM
        1. Smaller in size
        1. Less expensive
        1. Less power consumption

# Stack and Heap

## Key Difference Between Stack and Heap Memory
* Stack is a linear data structure whereas Heap is a hierarchical data structure.
* Stack memory will never become fragmented whereas Heap memory can become fragmented as blocks of memory are first allocated and then freed.
* Stack accesses local variables only while Heap allows you to access variables globally.
* Stack variables can’t be resized whereas Heap variables can be resized.
* Stack memory is allocated in a contiguous block whereas Heap memory is allocated in any random order.
* Stack doesn’t require to de-allocate variables whereas in Heap de-allocation is needed.
* Stack allocation and deallocation are done by compiler instructions whereas Heap allocation and deallocation is done by the programmer.

## What is a Stack?
A stack is a special area of computer’s memory which stores temporary variables created by a function. In stack, variables are declared, stored and initialized during runtime.

It is a temporary storage memory. When the computing task is complete, the memory of the variable will be automatically erased. The stack section mostly contains methods, local variable, and reference variables.

## What is Heap?
The heap is a memory used by programming languages to store global variables. By default, all global variable are stored in heap memory space. It supports Dynamic memory allocation.

The heap is not managed automatically for you and is not as tightly managed by the CPU. It is more like a free-floating region of memory.

## Key Differences between Stack and Heap



```
        Parameter	                Stack	                        Heap
    ----------------------------------------------------------------------
    Type of data            A stack is a linear             Heap is a hierarchical 
    structures	            data structure.                 data structure.

    Access speed	        High-speed access	            Slower compared to stack

    Space management	    Space managed efficiently       Heap Space not used as efficiently. 
                            by OS so memory will never      Memory can become fragmented as  
                            become fragmented.	            blocks of memory first              
                                                            allocated and then freed.
    
    Access	                Local variables only	        It allows you to access variables   
                                                            globally.

    Limit of space size	    Limit on stack size             Does not have a specific 
                            dependent on OS.	            limit on memory size.

    Resize	                Variables cannot be resized	    Variables can be resized.

    Memory Allocation	    Memory is allocated in a 	    Memory is allocated in any random   
                            contiguous block.               order.

    Allocation and	        Automatically done by 	        It is manually done 
    Deallocation            compiler instructions.          by the programmer.

    Deallocation	        Does not require to 	        Explicit de-allocation is 
                            de-allocate variables.          needed.

    Cost	                Less	                        More

    Implementation	        A stack can be implemented      Heap can be implemented 
                            in 3 ways simple array based,   using array and trees.
                            using dynamic memory, 
                            and Linked list based.	


    Main Issue	            Shortage of memory	            Memory fragmentation

    Locality of 	        Automatic compile               Adequate
    reference               time instructions.	

    Flexibility	            Fixed size	                    Resizing is possible

    Access time	            Faster	                        Slower
```

## Differences between Stack and Heap

Stack and a Heap ?

Stack is used for static memory allocation and Heap for dynamic memory allocation, both stored in the computer's RAM .

Variables allocated on the stack are stored directly to the memory and access to this memory is very fast, and it's allocation is dealt with when the program is compiled. When a function or a method calls another function which in turns calls another function etc., the execution of all those functions remains suspended until the very last function returns its value. The stack is always reserved in a LIFO order, the most recently reserved block is always the next block to be freed. This makes it really simple to keep track of the stack, freeing a block from the stack is nothing more than adjusting one pointer.

Variables allocated on the heap have their memory allocated at run time and accessing this memory is a bit slower, but the heap size is only limited by the size of virtual memory . Element of the heap have no dependencies with each other and can always be accessed randomly at any time. You can allocate a block at any time and free it at any time. This makes it much more complex to keep track of which parts of the heap are allocated or free at any given time.

http://net-informations.com/faq/net/img/stack-heap.png

You can use the stack if you know exactly how much data you need to allocate before compile time and it is not too big. You can use heap if you don't know exactly how much data you will need at runtime or if you need to allocate a lot of data.

In a multi-threaded situation each thread will have its own completely independent stack but they will share the heap. Stack is thread specific and Heap is application specific. The stack is important to consider in exception handling and thread executions.

## Memory Layout of C Programs
* A typical memory representation of a C program consists of the following sections.
    1. Text segment  (i.e. instructions)
    1. Initialized data segment 
    1. Uninitialized data segment  (bss)
    1. Heap 
    1. Stack 

## Stack vs Heap Memory Allocation

