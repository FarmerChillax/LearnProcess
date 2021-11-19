# -*- coding: utf-8 -*-
'''
    :file: schemas.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 22:15:44
'''

from apiflask import Schema 
from marshmallow.fields import Boolean, String

class ScanInSchema(Schema):
    host = String(required=True)
    all_port = Boolean(missing=False, default=False)


# class ScanOutSchema(Schema):
    
