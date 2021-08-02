# -*- coding: utf-8 -*-
'''
    :file: models.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/01/31 15:24:36
'''
from datetime import datetime
from werkzeug.security import generate_password_hash, check_password_hash

from flask_login import UserMixin

from app import db

class User(db.Model, UserMixin):
    id = db.Column(db.Integer,  primary_key=True)
    username = db.Column(db.String(20))
    password_hash = db.Column(db.String(128))
    last_login = db.Column(db.DateTime, default=datetime.utcnow)

    def set_password(self, password):
        self.password_hash = generate_password_hash(password=password)

    def validate_password(self, password):
        return check_password_hash(self.password_hash, password)

