# -*- coding: utf-8 -*-
'''
    :file: main.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/11/24 12:23:01
'''
import pyquery
import requests
from pyquery import PyQuery as pq

url = 'http://glidedsky.com/level/web/crawler-basic-1'



with open("index.html", 'r', encoding='utf-8') as f:
    html = f.read()


doc = pq(html)
items = doc('#app > main > div.container > div > div > div').children('.col-md-1')


sum = 0
for item in items.text().split(' '):
    sum += int(item)

print(sum)
