# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/06/04 15:12:06
'''
from flask import Flask
import os
import click
# import set
from app.settings import config
# imoprt extension
from app.extensions import db, cors
# import bp
from app.views.users import user_bp
# other
from flask_bootstrap import Bootstrap

def create_app(config_name = None):
    app = Flask(__name__)
    if config_name is None:
        config_name = os.getenv('FLASK_CONFIG', 'development')

    app.config.from_object(config[config_name])
    print(app.config['SECRET_KEY'])
    Bootstrap(app)
    register_extensions(app)
    register_blueprints(app)
    register_commands(app)
    return app

def register_extensions(app:Flask):
    db.init_app(app)
    cors.init_app(app)

def register_blueprints(app:Flask):
    app.register_blueprint(user_bp, url_prefix="/users")

def register_commands(app:Flask):
    from app.models import User
    @app.cli.command()
    def initdb():
        db.create_all()
        admin:User = User()
        admin.username = "farmer"
        admin.set_password("admin")
        users = []
        for i in range(100):
            user = User()
            user.username = f'{i}'
            user.set_password(f'{i}')
            users.append(user)
            # db.session.add(user)

        db.session.add_all(users)
        db.session.add(admin)
        db.session.commit()
        
        click.echo("initdb.")
