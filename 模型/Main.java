import java.util.LinkedList;

public class Main {
    public static void main(String[] args) throws InterruptedException {
        MessageQueue mq = new MessageQueue();


        // 开100个线程生产消息
        for (int i = 0; i < 100; i++) {
            int finalI = i;
            new Thread(() -> {
                mq.product("" + finalI);
                System.out.println("生产消息：" + finalI);
            }).start();
        }

        // 开一个线程消费消息
        new Thread(() -> {
           while (true) {
               String msg = mq.consume();
               System.out.println("消费消息：" + msg);
               if (msg == null) {
                   System.out.println("队列关闭");
                   break;
               }
           }
        }).start();

        // 等待消费
        Thread.sleep(100);
        mq.stop();
        System.out.println("主线程结束");
    }
}

/**
 * 使用锁实现的消息队列
 */
class MessageQueue {
    // 队列关闭标志
    private volatile boolean stop = false;
    // 使用链表作为消息容器
    private final LinkedList<String> queue = new LinkedList<>();
    /**
     * 生产消息
     * @param msg   消息内容
     */
    public void product(String msg) {
        // 获取链表的锁之后，向链表添加消息，最后尝试唤醒链表对象（如果在等待的话）
        synchronized (queue) {
            queue.add(msg);
            queue.notify();
        }
    }
    /**
     * 关闭消息队列，并唤醒消费者的等待（如果在等待的话）
     */
    public void stop() {
        synchronized (queue) {
            stop = true;
            queue.notify();
        }
    }
    /**
     * 消费消息，若无消息将会一直阻塞，直到有消息或队列关闭，若消息队列已关闭则不再阻塞且返回null
     * @return 拿到的消息
     */
    public String consume() {
        // 获取链表的同步锁
        synchronized (queue) {
            // 如果队列未关闭 就一直尝试获取一个消息
            while (!stop) {
                // 若链表长度为空，则临时释放掉链表的锁并使当前线程进入等待状态
                while (!stop && queue.size() == 0) {
                    try {
                        // 进入等待，直到queue被执行notify()或其他任意Object类的notifyAll()方法
                        queue.wait();
                    } catch (InterruptedException ignore) {
                        if (stop) return null;
                    }
                    // 被唤醒（同时重新获得锁），检查队列是否为已关闭，是的话就返回null，不是的话重新进入等待状态
                    if (stop) return null;
                }
                // 移除掉表头元素并返回
                return queue.remove();
            }
            return null;
        }
    }
}
