# -*- coding: utf-8 -*-
'''
    :file: views.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/12 17:31:19
'''
from flask import json, jsonify
from app import app

@app.route("/")
def index():
    return jsonify({
        "point": "index"
    })

