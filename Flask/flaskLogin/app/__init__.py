# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/01/31 15:24:16
'''
from flask import Flask

from flask_sqlalchemy import SQLAlchemy
from flask_login import LoginManager

app = Flask(__name__)
app.config.from_pyfile('settings.py')

db = SQLAlchemy(app)
login_manager = LoginManager(app)

@login_manager.user_loader
def load_user(user_id):
    from app.models import User
    user = User.query.get(int(user_id))
    return user

@login_manager.request_loader
def load_user_from_request(request):
    from app.models import User
    api_key = request.args.get('api_key')
    if api_key:
        user = User.qeury.filter_by(api_key=api_key).first()
    if user:
        pass
    return user
    


from app import commands, models
from app.views import login_bp

app.register_blueprint(login_bp)

login_manager.login_view = 'login_bp.login'