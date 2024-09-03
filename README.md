# golang八股文（面试题）


# 一、golang基础篇

## 1.5 Go 语言中的深拷贝和浅拷贝？

在 Go 语言中，当我们谈论深拷贝和浅拷贝时，其实是在讨论如何复制数据结构，如数组、切片、或结构体。理解这两者的区别对写出高效且安全的代码至关重要。

浅拷贝

**定义**：  
浅拷贝是指仅复制对象的引用（指针），而不复制实际的数据。这意味着拷贝后的新对象与原对象共享同一块内存区域。

**特性**：  
使用浅拷贝后，两个对象指向同一个底层数据。因此，如果我修改了其中一个对象的内容，另一个对象也会受到影响。

**例子**：

```go
original := []int{1, 2, 3}
shallowCopy := original // 浅拷贝

shallowCopy[0] = 10
fmt.Println("Original:", original)       // [10, 2, 3]
fmt.Println("Shallow Copy:", shallowCopy) // [10, 2, 3]
```

在这个例子中，我创建了一个切片的浅拷贝 `shallowCopy`，它实际上指向了 `original` 的相同底层数组。因此，对 `shallowCopy` 的修改会直接反映在 `original` 上。

深拷贝

**定义**：  
深拷贝则是完整复制对象及其所有引用的数据，确保新对象与原对象完全独立。

**特性**：  
深拷贝后的两个对象完全独立存在，互不影响。修改其中一个对象不会影响到另一个对象。

**例子**：

```go
original := []int{1, 2, 3}

// 手动实现深拷贝
deepCopy := make([]int, len(original))
copy(deepCopy, original)

deepCopy[0] = 10
fmt.Println("Original:", original)    // [1, 2, 3]
fmt.Println("Deep Copy:", deepCopy)   // [10, 2, 3]
```

在这里，我手动实现了一个深拷贝 `deepCopy`。通过 `copy` 函数，我将 `original` 的内容复制到 `deepCopy`，并确保它们之间是独立的。因此，修改 `deepCopy` 不会影响到 `original`。

何时使用

- **浅拷贝**：如果我希望多个对象共享相同的数据，或在处理大数据时希望提高性能，浅拷贝是个好选择。但要小心，它可能导致数据被无意修改，特别是在并发环境下。

  - ```go
    对于引用类型来说，你的每一次拷贝，Go 不会申请新的内存空间，而是使用它的指针，两个变量名其实都指向同一块内存空间，改变其中一个变量，会直接影响另一个变量。
    ```

- **深拷贝**：当我需要确保对象的独立性，不希望不同对象之间的修改互相影响时，深拷贝就显得尤为重要。尤其是在并发编程中，为每个 goroutine 提供独立的数据副本是非常必要的。

总结

浅拷贝和深拷贝各有其适用场景。选择哪种拷贝方式，取决于我的具体需求和应用场景。在需要共享数据的场景下，我倾向于使用浅拷贝；而在需要独立处理的数据结构中，我更愿意使用深拷贝来避免潜在的错误。

## 1.6 Go 是值传递，还是引用传递、指针传递？

在 Go 语言中，函数参数的传递方式可以归类为**值传递**。这意味着在函数调用时，会将参数的副本传递给函数，而不是直接传递原始值或引用。然而，由于 Go 支持指针，可以通过指针传递来间接修改原始值，这往往引发关于“引用传递”或“指针传递”的讨论。让我详细解释一下：

1. **值传递**

- **定义**：值传递是将变量的副本传递给函数。在函数内部修改该副本，不会影响到原始变量。

**例子**：

```go
func modifyValue(x int) {
    x = 10
}

func main() {
    a := 5
    modifyValue(a)
    fmt.Println(a) // 输出: 5
}
```

在这个例子中，我将变量 `a` 的副本传递给了 `modifyValue` 函数。在函数内部，虽然我修改了 `x` 的值，但这并不会影响到原始变量 `a`。

2. **指针传递**

- **定义**：通过传递指针（即内存地址），函数可以直接修改原始变量的值。这种传递方式仍然是值传递，因为传递的是指针的副本，但由于指针指向了原始数据，函数可以通过指针操作原始数据。

**例子**：

```go
func modifyPointer(x *int) {
    *x = 10
}

func main() {
    a := 5
    modifyPointer(&a)
    fmt.Println(a) // 输出: 10
}
```

在这个例子中，我传递了 `a` 的指针给 `modifyPointer` 函数。在函数内部，通过指针修改了原始变量 `a` 的值，所以 `a` 的值在函数调用后变成了 `10`。

3. **引用传递（Go 不支持）**

- **定义**：引用传递是指将变量的引用直接传递给函数，函数内部操作引用即操作原始变量。Go 语言并不直接支持引用传递，而是通过值传递和指针传递间接实现类似的效果。

### 总结

- **值传递**：Go 中所有的参数传递本质上都是值传递，即使是传递指针或切片也是如此。只不过传递的是指针的副本或切片的引用，而不是它们所指向的数据的副本。
- **指针传递**：通过传递指针，可以在函数内部修改原始变量的值，这种方式仍然是值传递，只不过值是一个指针。









































































1. go语言有哪些优点、特性？
语法简便，容易上手。

支持高并发，go有独特的协程概念，一般语言最小的执行单位是线程，go语言支持多开协程，协程是用户态线程，协程的占用内存更少，协程只独有自己的栈等一些资源。其他都是共享线程的。调度时可以极大减少上下文切换。

2. go的包管理
golang在1.11版本之前用的是GOPATH，项目代码要放在GOPATH/src目录下才可运行。但是GOPATH存在弊端，没有版本控制概念，go get的外部库，自己都不知道下载的什么版本，再多人协作开发时无法确定大家的外部库版本是否一致，会出大问题。

golang1.11后推荐使用go mod 管理，go env 下设置 GO111MODULE = on 打开。支持依赖升降级。


3. make和new 的一些区别
new 可以初始化声明所有类型，返回的是一个指针类型。对应的是该类型的空值

make 只能用来初始化 map slice channel 的，返回的就是这个类型。对应的是该类型的零值

new的用法不常见，一般使用 var 或 :=


4. channel 是可以被 close 的，之后还可以读写吗？
不管是有无缓冲channel。关闭后都可以读，如果缓冲区里还有数据，先将缓冲区数据读完后，在读读的是channel 传输类型的 零值。

不管是有无缓冲channel，关闭后都不可以写，写会panic


5. channel有什么用
channel 用来做协程之间的通信。go语言提倡不要用共享内存来通信，要用通信的方式来共享内存。channel就是利用通信的方式实现对资源的访问。

无缓冲channel可以用来做协程间的同步

有缓冲channel可以用来做消息队列

6. channel的底层结构？接收、发送消息的过程？
channel底层是一个结构体，里面包含一个环形缓冲区，还有接收等待队列和发送等待队列。分别存放等待写消息的协程队列和读消息的协程队列，还有一个互斥锁，用来保证channel的线程安全，防止多个协程并发读写可能导致的问题。

当一个协程向channel写数据，如果是无缓冲通道，会将其加入到消息等待发送队列中，等待一个协程向channel读数据

读数据时，会先看缓冲区，如果没有再唤醒消息等待发送队列的第一个协程拿到信息，如果没有加入到消息等待接收队列，等待一个协程向channel写数据

7. channel 并发安全吗?
channel并发安全，底层结构体中有一个互斥锁，在并发环境下可以自动实现加锁解锁操作。


8. 怎么样输出一个有序的map
用slice存放map的key值，对slice进行排序，然后按照slice顺序查找map key value

用golang的数据结构 list来实现，list是一个链表。map[key]value value存放链表节点，写入map的时候同时按照顺序写入list中，顺序遍历时，只需要遍历链表即可。可以达到有序效果。


9. map在传参时的类型
golang传参都是值传递，只是对于map chan 等类型，他们在进行值拷贝的时候，调用的makechan 和 makemap 方法返回的都是指针类型的变量。导致操作的其实是同一个内存。

10. 可以直接对map取地址吗？
对map可以取址，但是对于map的key value不能取址，map不支持这种操作，当某种条件达到后，map会做增量扩容或者等量扩容操作。这时候每个key 的地址都可能发生变化。因此获取map key  value地址的操作是无意义的。

11. map底层结构
map底层是一个结构体，hmap，hmap中包含 元素的个数，桶的个数，以及指向桶数组的指针，每一个桶是一个bucket，bucket在go中用的是 bmap结构体，每个bucket可以存放 8个键值对，哈希值低八位相同的键存入bucket时会将高八位存储在tophash数组里。data 区域存放的key-value数据是按照 keykeykey valuevaluevalue存放的，overflow指针指向下一个bucket桶，通过链表将所有冲突的键连接起来。

12. 哪些数据类型不能作为map里面的key，哪些可以，有没有什么评判标准？
无法比较的数据类型都不能作为map里面的key。

基本数据类型都可作为key，指针也可以，因为指针比较的是地址。 数组也可以，结构体也可以

不可比较 map slice function

13. map是并发安全的吗？怎么实现并发安全？
map不是并发安全的，底层结构体没有用互斥锁进行加锁解锁操作。

如何实现并发安全，1.map+读写锁 （推荐）   2.用sync.Map


14. sync.Map如何解决并发问题？
sync.Map底层结构体中嵌入了一个互斥锁。当某个协程对map进行写操作时，会上写互斥锁，其他协程看到锁，于是会阻塞。当某个协程对map进行读操作时，会上读锁，但是读锁可以共享，所以其他协程也可以对map进行读操作。


15. gmp模型（局部饥饿、全局饥饿、全局缓存队列等）
最开始的时候，大多用线程池，开一定量的线程，当有工作任务到来时，会拿出一个线程处理，但当因为发生系统调用而阻塞时，线程池中可工作的线程就少了，线程池的性能就降低了。于是有了GMP模型。

G是协程，用户态线程

本地队列：每个P都有本地队列，本地队列中存放的是G

全局队列：当本地队列都满了，新来的G会优先加入本地队列中，本地队列满了会把本地队列的一半加入到全局队列

M是工作线程，用来处理协程的

P是处理器，默认数量是内核数，M通过获取P来处理P中的G

15. go有了协程之后，那它的线程是怎么调度的？（GMP模型）M最多多少个？
有调度策略：线程正常情况下会和P绑定，处理P中的本地队列G，但当处理G时，由于系统调用导致阻塞时，会触发hand off，M会和 P 解绑，把P转移给其他的空闲M，如果没有就创建一个新的M处理。如果M将P中的队列G处理完了，就会从其他P的本地队列中偷取一半协程来运行，如果其他都没有，就会从全局队列中选择一批来处理。

M最多10000个，在初始化运行时会设置最大值。但因为M需要获取到P才能处理G，P的数量是很有限的，所以M不会很多。

16. 线程数量和什么有关系？线程数量是无限大的吗？
线程数量和程序运行时，设置的最大M有关，默认是10000，并且跟P有关系，因为M必须要拿到P才能处理G。而且如果线程处理的阻塞的G完成后，可能会被销毁。不能无限大


16. slice深度拷贝
copy方法时，或者append触发扩容机制时都是深度拷贝。拷贝的是数据本身，新创建的对象和源对象不共享内存，会另开辟一个新的内存地址，值修改时不会影响源对象值。

17. slice和数组有什么区别？
slice 结构体中包含了 一个指向底层数组的指针还有slice的长度及容量

数组无法扩容，切片可以动态追加

数组声明时需要指定长度，切片不需要

数组是值类型的，切片是引用类型的


18. slice如何扩容？
1.18版本之前，当切片容量1024之前，每次扩容都是原来数组的二倍，当切片容量到1024之后，每次扩容都是原数组1.25倍

1.18版本，觉得扩容比例一下从2变成1.25有些不平滑，于是调整扩容机制，在容量小于256时，是2倍扩容，当容量大于256时，增加 （原来容量+3*256）/4。逐渐接近1.25倍，更平滑


19. interface{}怎么用
interface{}可以用来实现多态和符合设计模式的原则 依赖倒转 面向接口编程。 写一个接口，里面有要实现的方法，定义一个结构体，实现接口中的方法，然后初始化时，声明一个接口对象并指向具体的结构体，这样实现了面向接口编程，解耦。

也可以用来做强制类型转换，断言。interface.(type)


20. 说一下go的继承
go中的继承是通过结构体嵌套实现的，子结构体嵌套父结构体。初始化子结构体后，就可以用父结构体的方法。


21. go中哪些变量是不能比较的
map slice function 


22. golang强类型，弱类型？
golang是强类型语言，不会隐式的进行数据类型的转换。在速度上可能略逊弱类型语言，但是严谨性又避免了不必要的错误


23. golang并发控制？
可以用 Sync.WaitGroup实现  WaitGroup有三个方法，Add,Done,Wait。用来控制计数器数量。

可以用channel实现，goroutine之间通过channel通信，如果多个goroutine都写入channel，这时候通过缓冲区就可以实现并发控制，当缓冲区满时，后面的goroutine的写操作会被阻塞。


24. 一个go-routine最小占多大内存空间？
2kb


25. context类型有哪些？Context的作用是什么？context如何实现cancel的？
context主要用来在协程间传递关闭信号、信息的。可以控制多级协程

空context、cancelCtx、valueCtx、timerCtx。cancelCtx可以用来做协程的断开操作。




27. 正常模式和饥饿模式？
针对于互斥锁的，Mutex。正常模式下请求锁的协程要按照先入先出顺序排队，以此被唤醒，唤醒后还要与新请求锁的协程进行竞争，因为新请求的协程有优势，他们正在CPU上运行或者数量比较多。新唤醒的协程很难获取到锁，于是又会加到队列头部，如果一个等待的协程超过1ms仍未获取到锁，就会进入到饥饿模式。

饥饿模式下，互斥锁所有权会直接从解锁的协程转移到队首的协程，并且新到达的协程不会尝试获取锁而是加到队列尾部。如果一个等待协程获取到锁并且满足下面两个条件只一，就会回到正常模式。是队列的最后一个协程，等待时间小于1ms。

正常模式有更好的性能，饥饿模式可以避免尾部延迟这种情况


28. 用Go实现一个死锁
死锁存在四个条件，互斥，占有且等待，循环等待，不可强占用

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var x sync.WaitGroup
	x.Add(2)
	go func() {
		ch2 <- 1
		fmt.Println(<-ch1)
		x.Done()
	}()
	go func() {
        ch1 <- 1
		fmt.Println(<-ch2)
		x.Done()
	}()
	x.Wait()
}


29. gc的了解
gc是垃圾回收机制，为了对堆栈上使用完的对象及时回收。防止内存溢出。最初1.3版本之前使用的是标记清除法，GC时会暂停程序进行标记清楚操作。STW时间很长。为了优化STW，1.3版本后缩短了一些STW时间，回收对象的操作放在STW时间之外。1.5版本时为了降低STW的时间，使用三色标记法+插入/删除屏障实现。三色标记法是在GC开始时，从根节点开始遍历所有可达对象，第一次遍历到的对象置灰，第二次遍历到的对象置黑。当没有灰色对象是开始回收对象。但是这样存在问题，当新创建了一个白色对象时，被一个黑色对象引用了，此时会造成刚创建的对象也被回收。还有当删除一个对象时，但是后面的对象还想要使用，但是断掉后后面的对象也会被回收。针对这两种情况提出了两个设计原则，强三色不变式和弱三色不变式，强三色强调黑色对象不能引用白色对象。弱三色不变式强调当黑色对象引用白色对象是，要保证还有一个灰色对象引用白色对象。针对这两个原则，产出了插入屏障和删除屏障，插入屏障实现的是强三色不变式，当黑色对象引用白色对象时，会将白色对象变为灰色对象。删除屏障实现的是弱三色不变式，当删除灰色后面的白色对象，会把删除的白色对象置灰。当扫描完没有灰色对象后，会将栈对象置白，启用STW，从新扫描一遍栈上的可达节点，最终删掉白色对象。插入屏障仍然需要re-scan栈上的节点。删除屏障，删除效率低，这次要删除的对象要等下次才能删掉。于是1.8版本后，采用三色标记法+混合写屏障。混合写屏障结合了插入和删除屏障的优点，栈上的对象不启用GC，扫描时会将所有栈上对象和新建对象置黑，针对堆上对象，所有创建和删除的对象都置灰。体现的是变形的弱三色不变式。1.8版本极大减少了STW时间，但并不是完全没有，因为在GC三色标记法之前还要STW，开启辅助GC和写屏障，统计根对象的任务数量等。


30. 什么时候会触发 golang GC 呢？
手动触发：调用runtime.GC来手动触发

定期触发：最长两分钟触发一次GC

内存达到一定值：每当内存扩大一倍时启用GC


31. 内存泄漏
创建的资源没有正常释放造成内存一直占用。也无法被GC回收。内存泄漏一般是程序员申请资源后没有手动去释放资源，关闭通道，解锁等。

通过pprof工具排查

32. go里面声明一个变量，它是放在栈上还是堆上？
内存逃逸，局部变量从栈上逃逸到堆上发生了内存逃逸。 

对于指针类型，因为分配内存时不知道是否存在外部引用，于是就将内存分配在堆中，防止函数结束后局部变量被回收。

如果数据量太大，栈放不下就会逃逸到堆中

interface类型上调用方法，编译时不知道是如何实现的，就会放到堆中。


33. 子goroutine的panic会不会被父g捕获
panic只能捕获本协程内部的错误，无法捕获子协程的错误

34. defer的先后顺序
defer满足栈的结构，先入后出。

35. defer什么情况下可以修改函数的返回值？
defer和return知识，return分为两步操作，第一步先将返回值赋给return，在准备返回之前要执行defer函数，执行完后才将结果返回。 如果定义了具名返回值，并且在defer时对这个变量进行了处理。那return 这个变量时就会修改。


36. golang 如何做超时控制？
1.timeAfter 这个函数会返回一个通道，并且过一段时间后会自动向通道发送一个数据。

2.context 的 withTimeout 和 withDeadline 方法，当到规定时间或超过某一时间后，会调用cancel方法来关闭channel。可以做超时控制


37. select 一般使用在什么场景
select一般与通道连用，用来监听多个通道的状态。

可以在多个通道中选择一个可用的操作来执行，如果没有一个通道准备就绪，select会执行defaule语句或者继续等待。

38.go程序运行顺序？
从main包开始执行，首先会导入引用的包，如果这个包里有init函数，会先执行这个包的init函数。

如果在main包定义了init函数，Go运行会在包导入后，全局变量初始化之前调用init函数。

调用之后会进行全局变量初始化操作，最终进入到main函数，是整个程序入口点。main函数执行后，程序会按照调用函数顺序依次执行其他函数。

原文链接：https://blog.csdn.net/qq_67503717/article/details/136386099































# Golang面试问题汇总:



1.说说为什么同样实现一个“hello world”，Go编译出来的程序一般会比C/C++要大？ 关键字：跨平台 静态链接编译 依赖库 自带runtime

2.说说channel的实现.（核心，拓展问题：通信常用手段，阻塞非阻塞，同步异步的区别，select／poll/epoll等等）

3.goruntine是怎么调度的？与进程，线程的关系。（核心，拓展问题，进程，线程，协程区别，死锁，操作系统等等）

4.如何理解“不要通过共享内存来通信，要通过通信来共享内存”这句话？（关于代码和设计）

关键字：高内聚低耦合 消息机制 channel 结合场景

5.说说排查Go问题的经历，都用到了什么工具，有什么看法？（经验，排查问题思路和能力，个人觉得排查问题重要的是有没有思路方法，不见得什么最有效）

6.Go目前的GC策略是什么？之前是什么？怎么改进的？（拓展问题：关于GC算法，内存分配等等）

## 关于 goroutine 的退出：



那么我们怎么来控制一个goroutine，使它可以结束自己的使命正常结束呢？其实很简单，同样我们使用`select + return`来实现这个功能。

```
func boring(msg string, quit chan bool) <-chan string { 
    c := make(chan string)
    go func() { 
        for i := 0; ; i++ {
        	select {
        	case c <- fmt.Sprintf("%s: %d", msg, i):
        		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)	
        	case <-quit:
        		return
        	}
        }
    }()
    return c
}
func main(){
	quit := make(chan bool)
    c := boring("Joe", quit)
    for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
    quit <- true
}
```



## Go语言对网络IO的优化



主要是从两方面下手的：

- a. 将标准库中的网络库全部封装为非阻塞形式，防止其阻塞底层的M并导致内核调度器切换上下文带来的系统开销。
- b. 运行时系统加入epoll机制(针对Linux系统)，

> 当某一个Goroutine在进行网络IO操作时，如果网络IO未就绪，就将其该Goroutine封装一下，放入epoll的等待队列中，当前G挂起，与其关联的M可以继续运行其他G。

> 当相应的网络IO就绪后，Go运行时系统会将等待网络IO就绪的G从epoll就绪队列中取出（主要在两个地方从epoll中获取已网络IO就绪的G列表，一是sysmon监控线程中，二是自旋的M中），再由调度器将它们像普通的G一样分配给各个M去执行。

Go语言将高性能网络IO的实现方式直接集成到了Go本身的运行时系统中，与Go的并发调度系统协同高效的工作，让开发人员可以简单，高效地进行网络编程。

## 1. 除了 mutex 以外还有那些方式安全读写共享变量？



A。封装一些线程读写安全的类型作为共享变量，如`Atomic`。

B，将共享变量的读写放到一个 goroutine 中，其它 goroutine 通过 channel 进行读写操作。

c,最常见的互斥了，共享变量要互斥访问:

```
    sema <- struct{}{} // acquire token
    balance = balance + amout
    <-sema // release token
```



## 2. 无缓冲 chan 的发送和接收是否同步?



当通道中有数据时，向通道中发送会阻塞，当channel中无数据是，从channel中接收会阻塞。是否同步取决于使用者的操作。

同步的概念是：执行一个函数或者方法后，必须等到返回信息才能进行下一步操作。**从这方面理解的话，无缓冲的channel可以理解为同步的。因为发送必须等待channel中有数据，channel接收必须等待channel中无数据**

## 3. go语言的并发机制以及它所使用的CSP并发模型．



golang的并发建立在 goroutines，channels，select等的基础设施之上

csp并发模型：CSP 描述这样一种并发模型：多个Process 使用一个 Channel 进行通信, 这个 Channel 连结的 Process 通常是匿名的，消息传递通常是同步的。

go 其实只用到了 csp 模型的一小部分，即理论中的 **Process/Channel（对应到语言中的 goroutine/channel）**：这两个并发原语之间没有从属关系， Process 可以订阅任意个 Channel，**Channel 也并不关心是哪个 Process 在利用它进行通信**；Process 围绕 Channel 进行读写，形成一套有序阻塞和可预测的并发模型。

## 4. golang 中常用的并发模式？



- 通过sync包中的WaitGroup实现并发控制
- Context上下文，实现并发控制

> 根据 Rob pike 在 Google I/O 2012上的[演讲](http://talks.golang.org/2012/concurrency.slide#55)，golang的并发建立在 goroutines，channels，select等的基础设施之上，总结起来主要是：（这些都是基于csp并发模型进行的）

- 生产者消费者模型（多对一）
- 发布订阅模型（多对多m:n）【在这个模型中，消息生产者成为发布者（publisher），而消息消费者则成为订阅者（subscriber），生产者和消费者是M:N的关系】

## nil slice和empty slice的区别



以下是错误的用法，==会报数组越界的错误==，因为只是声明了slice，却没有给实例化的对象，这一点如果是cpp的vector，便可以直接使用，但是golang 不行。

```
var slice []int
slice[1] = 0
```



此时**slice的值是nil**，这种情况可以用于需要返回slice的函数，当函数出现异常的时候，保证函数依然会有nil的返回值。

empty slice 是指slice不为nil，但是slice没有值，**slice的底层的空间是空的**，此时的定义如下：

```
slice := make([]int,0）//或者
slice := []int{}
```



## 5. JSON 标准库对 nil slice 和 空 slice 的处理是一致的吗？　



**==不一致==！！！**

```
package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Num []int
}


func main(){
	var arr1 []int
	arr2 := make([]int,0)
	one := A{"one",arr1}
	two := A{"two",arr2}

	one1,err := json.Marshal(one)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(one1))

	one2,err := json.Marshal(two)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(one2))
}

----------
$ go run json.go
{"Name":"one","Num":null}
{"Name":"two","Num":[]}
```



在编码上，**==nil== slice 返回的是（json）null，而空 slice 返回的是空数组**。

## 6. 协程，线程，进程的区别。



> 进程是系统资源分配的最小单位, 系统由一个个进程(程序)组成 一般情况下，包括文本区域（text region）、数据区域（data region）和堆栈（stack region）。

- 文本区域存储处理器执行的代码
- 数据区域存储变量和进程执行期间使用的动态分配的内存；
- 堆栈区域存储着活动过程调用的指令和本地变量。

通信问题: 由于进程间是隔离的,各自拥有自己的内存内存资源, 因此相对于线程比较安全, 所以不同进程之间的数据只能通过 IPC(Inter-Process Communication) 进行通信共享.

> 线程:线程属于进程,共享进程的内存地址空间。同时多线程是不安全的,**当一个线程崩溃了,会导致整个进程也崩溃了,即其他线程也挂了,** 但**多进程而不会**,一个进程挂了,另一个进程依然照样运行。

进程切换分3步:

- 切换页目录以使用新的地址空间
- 切换内核栈
- 切换硬件上下文

而线程切换只需要第2、3步,因此进程的切换代价比较大

> 协程:协程是属于线程的。协程程序是在线程里面跑的，因此协程又称微线程和纤程等**。协没有线程的上下文切换消耗**。协程的调度切换是用户(程序员)手动切换的,因此更加灵活,因此又叫**用户空间线程**.

## 7. 互斥锁，读写锁，死锁问题是怎么解决。



编码过程中加锁的地方记得释放锁。

`go tool vet`在编译阶段发现一些常规错误问题。 `go run -race`检测竞争态问题。

## 9. Data Race问题怎么解决？能不能不加锁解决这个问题？



- 避免对共享 variable 进行写操作

- 避免在多个 goroutine 中存取变量(使用channel)

- 使用锁来保护共享的 variable 来保证任意时刻只有一个 goroutine 对变量进行修改

- > `go run -race`检测竞争态问题。 能不能不加锁解决这个问题: 使用channel当作锁：`done <- struct{}{}`

## 10. 什么是channel，为什么它可以做到线程安全？



channel 的实现就是 队列 + 锁，源码分析可知。

1. goroutine 如何调度?
2. Golang GC 时会发生什么?
3. Golang 中 goroutine 的调度.
4. 并发编程概念是什么？
5. 负载均衡原理是什么
6. lvs相关
7. 微服务架构是什么样子的

## 18. 分布式锁实现原理，用过吗？



在分布式环境中某一个值或状态只能是唯一存在的，这是分布式锁实际的基本原理。 实现方式：

##### 基于数据库实现分布式锁



> 直接创建一张锁表，然后通过操作该表中的数据来实现了。当我们要锁住某个方法或资源时，我们就在该表中增加一条记录，想要释放锁的时候就删除这条记录。

###### 基于缓存实现分布式锁（redis，etcd）



> 基于缓存来实现在性能方面会表现的更好一点。而且很多缓存是可以集群部署的，可以解决单点问题

###### 基于Zookeeper实现分布式锁



> 每个客户端对某个方法加锁时，在zookeeper上的与该方法对应的指定节点的目录下，生成一个唯一的瞬时有序节点。 判断是否获取锁的方式很简单，只需要判断有序节点中序号最小的一个。 当释放锁的时候，只需将这个瞬时节点删除即可。同时，其可以避免服务宕机导致的锁无法释放，而产生的死锁问题。

## 19. etcd和redis怎么实现分布式锁



```
> set lock:codehole true ex 5 nx 
OK
... do something critical ...
> del lock:codehole
```



上面这个指令就是 setnx 和 expire 组合在一起的原子指令，它就是分布式锁的奥义所在。

#### 超时问题



Redis 的分布式锁不能解决超时问题，如果在加锁和释放锁之间的逻辑执行的太长，以至于超出了锁的超时限制，就会出现问题。

因为这时候第一个线程持有的锁过期了，临界区的逻辑还没有执行完，这个时候第二个线程就提前重新持有了这把锁，导致临界区代码不能得到严格的串行执行。

为了避免这个问题，**Redis分布式锁不要用于较长时间的任务。**

**更好一点的方案：为set指令的value参数设置为一个随机数，释放锁时先匹配随机数是否一致，然后再删除 key，这是为了确保当前线程占有的锁不会被其它线程释放，除非这个锁是过期了被服务器自动释放的**。

1. goroutine和channel的作用分别是什么
2. 怎么查看goroutine的数量
3. 说下go中的锁有哪些? 三种锁，读写锁，互斥锁，还有map的安全的锁
4. 读写锁或者互斥锁读的时候能写吗
5. 怎么限制goroutine的数量

## 28. channel是同步的还是异步的



同步的

## 29. 说一下异步和非阻塞的区别



> 同步与异步:同步和异步关注的是消息通信机制 (synchronous communication/ asynchronous communication)

> 阻塞与非阻塞 :阻塞和非阻塞关注的是程序在等待调用结果（消息，返回值）时的状态.

## 30. log包线程安全吗？



go 标准库的log包是线程安全的。

1. goroutine和线程的区别
2. 滑动窗口的概念以及应用
3. 怎么做弹性扩缩容，原理是什么
4. 让你设计一个web框架，你要怎么设计，说一下步骤

## 35. 说一下中间件原理



> 类似于 Python 的装饰器，这归功于golang将函数作为一直引用类型，函数在go的类型系统里也是一等公民。而 go 的http中间件实现主要是利用高阶函数来设计。

**go 原生的http服务器主要依赖两种数据类型 ServeMux（map + 锁），http.Handler。http服务器处理路由与响应方法用 `http.Handle()`,`http.Handle`用`ServeMux`路由分发调用 `http.Handler`。**

一个go函数，或者对象方法，只要签名是 `func(ResponseWriter, *Request)` ，那么它的签名就与 `http.HandlerFunc` 一致，可以被强制转换为 `http.HandlerFunc`类型。同时`http.HandlerFunc`类型实现了`http.Handler`接口：

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request){
    f(w, r)
}
```



所以我们就可以将我们的函数或者方法包装成 `http.Handler`,然后在返回一个新的Handler，这就是中间件的实现过程。

## 36. 怎么设计orm，让你写你要怎么写



1. 反射获得各个字段的名称和值
2. 拼接sql，执行sql
3. 判断是否有自增字段，如果有，则设置自增字段的值。

## 37. epoll原理



> **epoll是一种I/O事件通知机制**

#### I/O事件



**I/O**: 输入输出(input/output)，输入输出的对象可以是 文件(file), 网络(socket), 进程之间的管道(pipe), 在linux系统中，都用文件描述符(fd)来表示

**事件**:

- 可读事件，当文件描述符关联的内核读缓冲区可读，则触发可读事件什么是可读呢？就是内核缓冲区非空，有数据可以读取
- 可写事件, 当文件描述符关联的内核写缓冲区可写，则触发可写事件什么是可写呢？就是内核缓冲区不满，有空闲空间可以写入

#### 通知机制



通知机制，就是当事件发生的时候，去通知他 通知机制的反面，就是轮询机制

以上两点结合起来理解

> **epoll是一种当文件描述符的内核缓冲区非空的时候，发出可读信号进行通知，当写缓冲区不满的时候，发出可写信号通知的机制**

1. 用过原生的http包吗？
2. 一个非常大的数组，让其中两个数想加等于1000怎么算

```
func twoSum(nums []int, target int) []int {
    var res []int
	co := make(map[int]int，len(nums))
	for i, v := range nums {
		co[v] = i
	}
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if value, ok := co[t]; ok && value != i {
			res = append(res, i)
			res = append(res, value)
			break
		}
	}
	return res
}
```



1. 各个系统出问题怎么监控报警

## 41. 常用测试工具，压测工具，方法



goconvey,vegeta

## 42. 复杂的单元测试怎么测试，比如有外部接口mysql接口的情况【goconvey】



其实很重要的一个部分就是被测代码本身是容易被测的，也就是说在设计和编写代码的时候就应该先想到相好如何单元测试，甚至有人提出**可以先写单元测试，再写具体被测代码。**因为一个接口(或者称为单元)在被设计好后，它实现就确定了，实际效果也确定了。

此外还可以通过Docker起一个mysql服务，运行create.sql和insert.sql,然后进行模拟测试。

1. redis集群，哨兵，持久化，事务
2. 高可用软件是什么
3. 怎么搞一个并发服务程序
4. 讲解一下你做过的项目，然后找问题问实现细节
5. mysql事务说一下
6. 怎么做一个自动化配置平台系统
7. grpc遵循什么协议
8. grpc内部原理是什么

## 52. http2的特点是什么，与http1.1的对比



```
| HTTP1.1                    | HTTP2       | QUIC                        |
| -------------------------- | ----------- | --------------------------- |
| 持久连接                       | 二进制分帧       | 基于UDP的多路传输（单连接下）            |
| 请求管道化                      | 多路复用（或连接共享） | 极低的等待时延（相比于TCP的三次握手）        |
| 增加缓存处理（新的字段如cache-control） | 头部压缩        | QUIC为 传输层 协议 ，成为更多应用层的高性能选择 |
| 增加Host字段、支持断点传输等（把文件分成几部分） | 服务器推送       |     
```
