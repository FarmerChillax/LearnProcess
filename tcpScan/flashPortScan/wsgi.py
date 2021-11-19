# -*- coding: utf-8 -*-
'''
    :file: wsgi.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 17:06:33
'''
import os
from dotenv import load_dotenv

dotenv_path = os.path.join(os.getcwd(), '.env')
if os.path.exists(dotenv_path):
    load_dotenv(dotenv_path)

from portScan import app, celery

