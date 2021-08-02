from app.models import Author, Article, Writer, Book, Citizen, City, Country, Capital
from app import app, db


@app.shell_context_processor
def make_shell_context():
    return dict(db=db, Author=Author, Article=Article, 
                Writer=Writer, Book=Book,
                Citizen=Citizen, City=City,
                Country=Country, Capital=Capital)

