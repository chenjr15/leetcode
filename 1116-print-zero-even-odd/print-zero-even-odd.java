class ZeroEvenOdd {
    private int n;
    Semaphore zero = new Semaphore(1);
    Semaphore even = new Semaphore(0);
    Semaphore odd = new Semaphore(0);
    // private volatile int i = 0;
    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i =1;i<=n;i++){
            zero.acquire();
            // ++i;
            printNumber.accept(0);
            if((i&1)==0){
                even.release();
            }else{
                odd.release();
            }

        }

    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i =1;i<=n;i++){
            if((i&1)==0){
                even.acquire();
                // ++i;
                printNumber.accept(i);
                zero.release();
            }
        }
        
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {

        for (int i =1;i<=n;i++){
            if((i&1)==1){
                odd.acquire();
                // ++i;
                printNumber.accept(i);
                zero.release();
            }
        }
    }
}