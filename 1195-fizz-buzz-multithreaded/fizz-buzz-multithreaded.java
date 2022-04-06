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