import os

from flask import Flask
from flask_mail import Mail
from flask_sqlalchemy import SQLAlchemy

# print(os.getenv('MAIL_USERNAME'), os.getenv('MAIL_PASSWORD'))

app = Flask(__name__)
app.config.from_pyfile('settings.py')

db = SQLAlchemy(app)

mail = Mail(app)


from app import views, commands, models