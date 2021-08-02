import socket
from queue import Queue
from concurrent.futures import ThreadPoolExecutor
import time

star_time = time.time()

socket.setdefaulttimeout(5)

class BoundThreadPoolExecutor(ThreadPoolExecutor):
    def __init__(self, *args, **kwargs) -> None:
        super(BoundThreadPoolExecutor, self).__init__(*args, **kwargs)
        self._work_queue = Queue(2)

def get_server_status(host: str = "farmer.xyz", port: int = 8088):
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        server.connect((host, port))
        print('{0} port {1} is open'.format(host, port))
        return True
    except Exception as err:
        # print('{0} port {1} is not open'.format(host, port))
        return False
    finally:
        server.close()

with BoundThreadPoolExecutor(max_workers=300) as pool:
    for num in range(1, 255):
        host = "172.16.1." + str(num)
        print(host)
        for port in range(1, 8000):
            future1 = pool.submit(get_server_status, host, port)
            # print(pool._work_queue.qsize())

end_time = time.time()
runtime = end_time - star_time
print("runtime: ", runtime)
# pool = ThreadPoolExecutor(max_workers=400)
# for num in range(1, 255):
#     host = "172.16.1." + str(num)
#     for port in range(1, 65535):
#         future1 = pool.submit(get_server_status, host, port)
# pool.shutdown()

# def get_que():
#     t = que.get()
#     t.start()

# def put_que(host, port):
#     print("putting...")
#     t = Thread(target=get_server_status,args=(host, port))
#     que.put(t)

# get = Thread(target=get_que, name="Get Thread")
# get.start()


        # put_que(host=host, port=port)
        # put = Thread(target=put_que, args=[host, port], name="Put Thread")
        

    # get_server_status(host=host + str(num), port=5000)