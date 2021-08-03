# -*- coding: utf-8 -*-
'''
    :file: settings.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 01:50:57
'''
import sys
import os

basedir = os.path.abspath(os.path.dirname(os.path.dirname(__file__)))

WIN = sys.platform.startswith('win')
if WIN:
    prefix = 'sqlite:///'
else:
    prefix = 'sqlite:////'

# 基本配置
class BaseConfig(object):
    # 鉴权加密密钥, 比如：cookie
    SECRET_KEY = os.getenv('SECRET_KEY', 'dev key')
    # ORM框架配置
    SQLALCHEMY_TRACK_MODIFICATIONS = False

# 开发环境配置
class DevelopmentConfig(BaseConfig):
    SQLALCHEMY_ECHO = True
    # 开发环境使用sqlite作为开发数据库
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', prefix + os.path.join(basedir, 'data.db'))

# 生产环境配置
class ProductionConfig(BaseConfig):
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', prefix + os.path.join(basedir, 'data.db'))


# export配置
config = {
    'development': DevelopmentConfig,
    'production': ProductionConfig
}