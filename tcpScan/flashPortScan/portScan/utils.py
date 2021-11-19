# -*- coding: utf-8 -*-
'''
    :file: utils.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 22:02:36
'''

from typing import Any


def make_res(code:int=200, message:str="ok", data:Any=None):
    return {
        "code": code,
        "message": message,
        "data": data,
    }
