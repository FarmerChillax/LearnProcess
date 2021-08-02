import click

from app import app, db
# from app.models

@app.cli.command()
@click.option('--drop', is_flag=True, help='Create after drop.')
def initdb(drop):
    if drop:
        click.confirm("This operation will delet the database, do you want to continue?", abort=True)
        db.drop_all()
        click.echo('Drop tables.')
    db.create_all()
    click.echo('Init database.')

@app.cli.command()
def drop():
    db.drop_all()
    click.echo('Drop all database.')