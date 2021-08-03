# -*- coding: utf-8 -*-
'''
    :file: model.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 02:05:00
'''


from app.extensions import db
import sqlalchemy as sa
from werkzeug.security import generate_password_hash, check_password_hash

class User(db.Model):
    id = sa.Column(sa.Integer, primary_key=True)
    # 注册必备
    name = sa.Column(sa.String(20))
    email = sa.Column(sa.String(260))
    password_hash = sa.Column(sa.String(128))
    blog = sa.Column(sa.String(260))
    phone = sa.Column(sa.Integer)

    def get_password(self) -> str:
        return self.password_hash
    
    def set_password(self, password):
        self.password_hash = generate_password_hash(password)

    password = property(get_password, set_password)

    def validate_password(self, password):
        return check_password_hash(self.password_hash, password)