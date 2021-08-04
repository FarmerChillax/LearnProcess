# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 01:47:47
'''

from app.model import User
import os
from apiflask import APIFlask
# 导入配置
from app.settings import config
from app.extensions import db, cors

# 导入视图
from app.blueprints.user import user_bp

def create_app(config_name=None)->APIFlask:
    if config_name is None:
        config_name = os.getenv("FLASK_CONFIG", "development")
    
    app = APIFlask(__name__)
    app.config.from_object(config[config_name])

    register_extensions(app)
    register_blueprints(app)
    register_hooks(app)
    return app


def register_blueprints(app:APIFlask):
    app.register_blueprint(user_bp, url_prefix="/user")

def register_extensions(app:APIFlask):
    db.init_app(app) # ORM 数据库操作
    cors.init_app(app) # CORS 跨域

def register_hooks(app:APIFlask):
    @app.before_first_request
    def init_db():
        users = [{
            "name": "farmer",
            "email": "farmer@farmer233.top",
            "blog": "https://blog.farmer233.top",
            "phone": 1234567890,
            "password": "admin"
        },{
            "name": "admin",
            "email": "admin@farmer233.top",
            "blog": "https://admin.com",
            "phone": 1234567890,
            "password": "admin"
        },]
        db.drop_all()
        db.create_all()
        for user in users:
            item = User(**user)
            db.session.add(item)
        db.session.commit()
