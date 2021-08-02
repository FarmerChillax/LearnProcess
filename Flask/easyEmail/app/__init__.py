from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_mail import Mail

import os

# flask 对象
app = Flask(__name__)

# 导入配置
app.config.from_pyfile('settings.py')

# 数据库对象
db = SQLAlchemy(app)

print(os.getenv('MAIL_USERNAME'), os.getenv('MAIL_PASSWORD'))
app.config.update(  # 邮件配置
    MAIL_SERVER = os.getenv('MAIL_SERVER'),
    MAIL_PORT=80,
    MAIL_USE_TLS=False,
    MAIL_USERNAME = os.getenv('MAIL_USERNAME'),
    MAIL_PASSWORD= os.getenv('MAIL_PASSWORD'),
    MAIL_DEFAULT_SENDER=('Farmer', os.getenv('MAIL_USERNAME'))
    )

# 邮件对象
mail = Mail(app)

from app import shells, commands, models, views