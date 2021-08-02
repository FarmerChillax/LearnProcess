# -*- coding: utf-8 -*-
'''
    :file: main.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/18 17:10:24
'''
from flask import Flask, request
# from flask_redis import FlaskRedis
from flask_cache import Cache


app = Flask(__name__)
app.config.from_pyfile('settings.py')
cache_client = Cache(app, config={
    "CACHE_TYPE": 'redis',
    'CACHE_REDIS_HOST': '192.168.2.122',
    'CACHE_REDIS_PORT': 6379,
    'CACHE_REDIS_PASSWORD': 'farmer233',  
    'CACHE_REDIS_DB': 0
})

@app.route("/")
def index():
    try:
        path = request.args.get("path")
        # print(path, data)
        # return data\
        return ""
    except Exception as e:
        print(type(e), e)
        return str(e)

@cache_client.cached(timeout=30, key_prefix="test")
def index():
    print("test value")
    return "test value"