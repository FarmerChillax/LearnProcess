# -*- coding: utf-8 -*-
'''
    :file: demo1.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/03/21 00:38:50
'''
import asyncio

async def main():
    await asyncio.create_task(
        asyncio.sleep(0.1)
    )

async def test():
    await asyncio.sleep(0.1)

asyncio.run(test())