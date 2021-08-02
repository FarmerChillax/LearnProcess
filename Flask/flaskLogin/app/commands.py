# -*- coding: utf-8 -*-
'''
    :file: commands.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/01/31 15:24:25
'''
import click

from app import app, db

@app.cli.command()
def init():
    db.drop_all()
    db.create_all()
    from app.models import User
    user = User()
    user.username = 'admin'
    user.set_password('admin')
    db.session.add(user)
    db.session.commit()
    click.echo('Init db.')
