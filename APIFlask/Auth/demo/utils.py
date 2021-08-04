# -*- coding: utf-8 -*-
'''
    :file: utils.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 22:17:08
'''

from flask.globals import current_app
from itsdangerous import BadSignature, SignatureExpired, TimedJSONWebSignatureSerializer as Serializer

def generate_auth_token(payload:dict, expiration=3600 * 24 * 7):
    """生成token

    Args:
        payload (dict): token载体
        expiration (int, optional): token有效时间. Defaults to 3600*24*7.

    Returns:
        str: token令牌
    """
    s = Serializer(current_app.config["SECRET_KEY"], expires_in=expiration)
    return s.dumps(payload).decode()

