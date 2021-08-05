# -*- coding: utf-8 -*-
'''
    :file: settings.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 16:45:56
'''

protocol="redis"
redis_host = f'farmer233.asuscomm.com'
redis_password='farmer233'
redis_port = 6379
redis_db=3
backend_db=4


redis_url = f'{protocol}://:{redis_password}@{redis_host}:{redis_port}/{redis_db}'
backend_url=f'{protocol}://:{redis_password}@{redis_host}:{redis_port}/{backend_db}'
