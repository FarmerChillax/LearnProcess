from queue import Queue
from threading import Thread
from time import sleep

que = Queue(5)


def get_obj():
    print("get obj running...")
    while True:
        t = que.get()
        sleep(5)
        t.start()

def make_obj():
    print("make obj running...")
    i = 1
    while True:
        t = Thread(target=print, args=["get ", i])
        print(que.qsize())
        que.put(t)
        
        i += 1

if __name__ == "__main__":
    get = Thread(target=get_obj)
    get.start()
    make = Thread(target=make_obj)
    make.start()
