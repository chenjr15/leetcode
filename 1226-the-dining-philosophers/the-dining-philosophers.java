class DiningPhilosophers {
    Semaphore maxp = new Semaphore(4);
    Semaphore[] chopsticks = new Semaphore[]{
        new Semaphore(1),
        new Semaphore(1),
        new Semaphore(1),
        new Semaphore(1),
        new Semaphore(1),
    };
    public DiningPhilosophers() {
        
    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {
        maxp.acquire();
        chopsticks[philosopher].acquire();
        chopsticks[(philosopher+1)%5].acquire();
        pickLeftFork.run();
        pickRightFork.run();

        eat.run();

        putRightFork.run();
        putLeftFork.run();

        chopsticks[philosopher].release();
        chopsticks[(philosopher+1)%5].release();

        maxp.release();
        
    }
}