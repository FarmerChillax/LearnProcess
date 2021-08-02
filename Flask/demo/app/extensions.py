# -*- coding: utf-8 -*-
'''
    :file: extensions.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/06/04 15:17:09
'''
from flask_cors import CORS
from flask_httpauth import HTTPTokenAuth
from flask_sqlalchemy import SQLAlchemy
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
from flask.globals import g, current_app


cors = CORS()
auth = HTTPTokenAuth()
# orm
db = SQLAlchemy()


# 初始化token鉴权插件
@auth.verify_token
def verify_token(token):
    g.user = None
    g.username = None
    try:
        data = Serializer(current_app.config['SECRET_KEY']).loads(token)
    except Exception as e:
        return False
    if "id" in data:
        g.user = data["id"]
        g.username = data['username']
        return True
    return False

