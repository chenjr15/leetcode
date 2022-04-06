# 1195.交替打印字符串

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/fizz-buzz-multithreaded/](https://leetcode-cn.com/problems/fizz-buzz-multithreaded/)

## 题目描述

<p>编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：</p>

<ul>
	<li>如果这个数字可以被 3 整除，输出 "fizz"。</li>
	<li>如果这个数字可以被 5 整除，输出 "buzz"。</li>
	<li>如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。</li>
</ul>

<p>例如，当 <code>n = 15</code>，输出： <code>1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz</code>。</p>

<p>假设有这么一个类：</p>

<pre>
class FizzBuzz {
  public FizzBuzz(int n) { ... }               // constructor
  public void fizz(printFizz) { ... }          // only output "fizz"
  public void buzz(printBuzz) { ... }          // only output "buzz"
  public void fizzbuzz(printFizzBuzz) { ... }  // only output "fizzbuzz"
  public void number(printNumber) { ... }      // only output the numbers
}</pre>

<p>请你实现一个有四个线程的多线程版  <code>FizzBuzz</code>， 同一个 <code>FizzBuzz</code> 实例会被如下四个线程使用：</p>

<ol>
	<li>线程A将调用 <code>fizz()</code> 来判断是否能被 3 整除，如果可以，则输出 <code>fizz</code>。</li>
	<li>线程B将调用 <code>buzz()</code> 来判断是否能被 5 整除，如果可以，则输出 <code>buzz</code>。</li>
	<li>线程C将调用 <code>fizzbuzz()</code> 来判断是否同时能被 3 和 5 整除，如果可以，则输出 <code>fizzbuzz</code>。</li>
	<li>线程D将调用 <code>number()</code> 来实现输出既不能被 3 整除也不能被 5 整除的数字。</li>
</ol>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>本题已经提供了打印字符串的相关方法，如 <code>printFizz()</code> 等，具体方法名请参考答题模板中的注释部分。</li>
</ul>

<p> </p>

## 标签

 - Concurrency 

## 代码

```java
class FizzBuzz {
    private int n;
    // int i=1;
    CyclicBarrier barrier = new CyclicBarrier(4);

    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        
        // while(i<=n){
        for (int i=1;i<=n;++i){
            // synchronized (this) {
                if (i<=n &&(i %3 ==0 && i%5 != 0)){
                    printFizz.run();
                    // ++i;
                }
                try{
                    barrier.await();
                }catch (BrokenBarrierException e) {
                    e.printStackTrace();
                }
            // }
            
        }
        
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        for (int i=1;i<=n;++i){
        //   while(i<=n){
            // synchronized (this) {
                if (i<=n &&(i %3 !=0 && i%5 == 0)){
                    printBuzz.run();
                    // ++i;
                }
            // }
                 try{
                    barrier.await();
                }catch (BrokenBarrierException e) {
                    e.printStackTrace();
                }

            
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        //   while(i<=n){
        for (int i=1;i<=n;++i){
            // synchronized (this) {
                //  System.out.println(i);
                if (i<=n &&(i %3 ==0 && i%5 == 0)){
                    printFizzBuzz.run();
                    // ++i;
                }
                  try{
                    barrier.await();
                }catch (BrokenBarrierException e) {
                    e.printStackTrace();
                }
                
            // }
            
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        for (int i=1;i<=n;++i){
        //  while(i<=n){
            // synchronized (this) {
                if (i<=n &&(i %3 !=0 && i%5 != 0)){
                    printNumber.accept(i);
                    // ++i;
                }
                try{
                    barrier.await();
                }catch (BrokenBarrierException e) {
                    e.printStackTrace();
                }
                
            // }
            
        }
    }
}
```