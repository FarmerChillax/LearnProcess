# -*- coding: utf-8 -*-
'''
    :file: settings.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/10 16:16:52
'''

import os
import sys

from sqlalchemy.sql.expression import true

basedir = os.path.abspath(os.path.dirname(os.path.dirname(__file__)))

# SQLite URI compatible
WIN = sys.platform.startswith('win')
if WIN:
    prefix = 'sqlite:///'
else:
    prefix = 'sqlite:////'

# 基本配置
class BaseConfig(object):
    # 鉴权加密密钥（cookie）
    SECRET_KEY = os.getenv('SECRET_KEY', 'dev key')
    # ORM框架配置
    SQLALCHEMY_TRACK_MODIFICATIONS = False

# 开发环境配置
class DevelopmentConfig(BaseConfig):
    # 开发环境使用sqlite作为开发数据库
    # SQLALCHEMY_DATABASE_URI = prefix + os.path.join(basedir, 'data-dev.db')
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', prefix + os.path.join(basedir, 'data.db'))

# 测试配置
class TestingConfig(BaseConfig):
    TESTING = True
    WTF_CSRF_ENABLED = False
    SQLALCHEMY_DATABASE_URI = 'sqlite:///:memory:'  # in-memory database

# 生产环境配置
class ProductionConfig(BaseConfig):
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', prefix + os.path.join(basedir, 'data.db'))


# xport配置
config = {
    'development': DevelopmentConfig,
    'testing': TestingConfig,
    'production': ProductionConfig
}