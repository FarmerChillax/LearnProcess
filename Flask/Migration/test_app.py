from flask import Flask
from flask_sqlalchemy import SQLAlchemy

import click
import os

app = Flask(__name__)

db = SQLAlchemy(app)

app.config['SQLALCHEMY_DATABASES_URI'] = os.getenv("DATABASE_URI", 'sqlite:///' + os.path.join(app.root_path, 'data.db'))

@app.cli.command()
def init():
    db.create_all()
    click.echo("init success")

@app.shell_context_processor
def make_shell_context():
    return dict(db=db)