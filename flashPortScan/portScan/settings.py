# -*- coding: utf-8 -*-
'''
    :file: settings.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 16:45:56
'''
import os

protocol="redis"
redis_host = os.getenv("REDIS_HOST")
redis_password=os.getenv("REDIS_PASSWORD")
redis_port = os.getenv("REDIS_PORT", 6379)
redis_db=os.getenv('REDIS_DB', 0)


redis_url = f'{protocol}://:{redis_password}@{redis_host}:{redis_port}/{redis_db}'
backend_url=f'{protocol}://:{redis_password}@{redis_host}:{redis_port}/{redis_db}'
