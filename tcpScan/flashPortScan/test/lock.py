# -*- coding: utf-8 -*-
'''
    :file: lock.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 16:19:55
'''

import time
import asyncio

async def fnuc1(lock:asyncio.Lock):
    print("func1 wait lock")
    async with lock:
        print("func1 get lock, doing...")
        time.sleep(2)
    
    print("func1 finish! released lock!")


async def fnuc2(lock:asyncio.Lock):
    print("func2 wait lock")
    async with lock:
        print("func2 get lock, doing...")
        time.sleep(2)
    
    print("func2 finish! released lock!")


lock = asyncio.Lock()


loop = asyncio.get_event_loop()

tasks = [fnuc1(lock), fnuc2(lock),fnuc1(lock),fnuc2(lock),fnuc1(lock)]

loop.run_until_complete(asyncio.wait(tasks))

print("done!")
