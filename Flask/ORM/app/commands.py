import click

from app import app, db
from app.models import Author, Article

@app.cli.command()
def init():
    db.create_all()
    click.echo("init success")

@app.cli.command()
def drop():
    click.echo("drop it now ~!")
    db.drop_all()


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
