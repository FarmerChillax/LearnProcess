import socket
import threading

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

for num in range(0, 255):
    host = "172.16.90." + str(num)
    t = threading.Thread(target=get_server_status,args=(host, 80))
    t.start()
    # get_server_status(host=host + str(num), port=5000)