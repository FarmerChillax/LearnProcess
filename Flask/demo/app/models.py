# -*- coding: utf-8 -*-
'''
    :file: models.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/06/04 15:39:36
'''
from app.extensions import db
import sqlalchemy as sa
from werkzeug.security import generate_password_hash, check_password_hash
from itsdangerous import BadSignature, SignatureExpired, TimedJSONWebSignatureSerializer as Serializer
from flask import current_app

class User(db.Model):
    # id = db.Column(db.Integer, primary_key=True)
    # id = sa.Column(sa.Integer, primary_key=True)
    # __table__ = "xiaotao"
    id = sa.Column(sa.Integer, primary_key=True)
    username = sa.Column(sa.String(20))
    password_hash = sa.Column(sa.String(128))

    def set_password(self, password):
        self.password_hash = generate_password_hash(password)

    def validate_password(self, password):
        return check_password_hash(self.password_hash, password)

    # 登录成功返回鉴权令牌（生成token）
    def generate_auth_token(self, expiration=3600 * 24 * 7):
        s = Serializer(current_app.config["SECRET_KEY"], expires_in=expiration)
        return s.dumps({
            'id': self.id,
            'username': self.username
        }).decode()