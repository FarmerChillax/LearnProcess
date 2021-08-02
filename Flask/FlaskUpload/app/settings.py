# -*- coding: utf-8 -*-
'''
    :file: settings.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/12 17:29:17
'''
import sys
import os

from app import app

WIN = sys.platform.startswith('win')
if WIN:
    prefix = 'sqlite:///'
else:
    prefix = 'sqlite:////'

dev_db = prefix + os.path.join(os.path.dirname(app.root_path), 'data.db')

SECRET_KEY = "test key"

SQLALCHEMY_TRACK_MODIFICATIONS = False
SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URI', dev_db)

# upload setting
# UPLOADED_FILES_DEST = os.path.join(os.path.dirname(app.root_path), "fileUpload")

# UPLOADED_FILES_URL = "ftp://www.xiaotao2333.top"