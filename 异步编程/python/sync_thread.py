# -*- coding: utf-8 -*-
'''
    :file: sync_thread.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/03/16 21:21:29
'''
import threading
storage = threading.local()
storage.test = 1

print(storage.test)

class AnotherThread(threading.Thread):

    def run(self):
        storage.test = 2
        print(storage.test)

another = AnotherThread()
another.start()

print(storage.test)