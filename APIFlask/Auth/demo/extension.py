# -*- coding: utf-8 -*-
'''
    :file: extension.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 22:11:15
'''

from apiflask import HTTPTokenAuth
from flask.globals import g, current_app
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
from flask_cors import CORS

cors = CORS()
auth = HTTPTokenAuth()

# 初始化token鉴权插件
@auth.verify_token
def verify_token(token):
    g.user = None
    try:
        data = Serializer(current_app.config['SECRET_KEY']).loads(token)
    except Exception as e:
        return False
    if "username" in data:
        g.user = data["username"]
        return True
    return False
