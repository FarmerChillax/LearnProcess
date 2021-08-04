# -*- coding: utf-8 -*-
'''
    :file: schemas.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 22:05:33
'''

from apiflask import Schema
from marshmallow.fields import Url, String
from marshmallow.validate import Length


class UserLoginSchema(Schema):
    username = String(required=True, validate=Length(5, 10))
    password = String(required=True, validate=Length(5, 20))

class UserRegInSchema(UserLoginSchema):
    url = Url()