# -*- coding: utf-8 -*-
'''
    :file: commands.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/09 18:11:48
'''
import click

from app import app, db

@app.cli.command()
@click.option('--drop', is_flag=True, help='Create after drop.')
def initdb(drop):
    '''Initialize the database.'''
    if drop:
        click.confirm('This operation will delet the database, do you want to continue?', abort=True)
        db.drop_all()
        click.echo('Drop tables.')
    db.create_all()
    click.echo("init database.")
