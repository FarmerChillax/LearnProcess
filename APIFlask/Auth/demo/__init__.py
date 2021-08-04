# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 22:16:40
'''

from apiflask import APIFlask, auth_required, input
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
from app.extension import cors
from app.views import bp


def create_app():
    app = APIFlask(__name__)
    app.config['SECRET_KEY'] = "YOUR_SECRET_KEY"
    register_extensions(app)
    register_blueprints(app)
    return app

def register_extensions(app:APIFlask):
    cors.init_app(app)

def register_blueprints(app:APIFlask):
    app.register_blueprint(bp, url_prefix="/")
