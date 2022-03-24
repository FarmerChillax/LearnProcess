import java.util.Random;

public class DeadLock {
    private Chopstick left, right;
    private Random random;

    public DeadLock(Chopstick left, Chopstick right) {
        this.left = left;
        this.right = right;
        random = new Random();
    }

    public void run() {
        try {
            while(true) {
                Thread.sleep(random.nextInt(1000));
                synchronized(left) {
                    synchronized(right) {
                        Thread.sleep(random.nextInt(1000));
                    }
                }
            }
        } catch (InterruptedException e) {
            //TODO: handle exception
        }
    }
}
