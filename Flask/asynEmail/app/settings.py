import os

from app import app

prefix = 'sqlite:///'

dev_db = prefix + os.path.join(os.path.dirname(app.root_path), 'data.db')

# 数据库设置
# 关闭警告提示
SQLALCHEMY_TRACK_MODIFICATIONS = False
# 设置数据库
SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URI', dev_db)

# 邮件配置
MAIL_SERVER = os.getenv('MAIL_SERVER')
MAIL_PORT=80
MAIL_USE_TLS=False
MAIL_USERNAME = os.getenv('MAIL_USERNAME')
MAIL_PASSWORD= os.getenv('MAIL_PASSWORD')
MAIL_DEFAULT_SENDER=('Farmer', os.getenv('MAIL_USERNAME'))