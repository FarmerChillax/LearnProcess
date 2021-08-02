# -*- coding: utf-8 -*-
'''
    :file: models.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/12 17:34:07
'''
from app import db

import SQLAlchemy as sa

class User(db.Model):
    id = sa.Column(sa.Integer, primary_key=True)
    path = sa.Column(sa.String(128))