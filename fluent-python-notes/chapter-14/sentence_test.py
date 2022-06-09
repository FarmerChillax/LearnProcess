# -*- coding: utf-8 -*-
'''
    :file: sentence_test.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/06/08 12:56:11
'''

from sentence_gen import Sentence

s = Sentence('"The time has come," the Walrus said.')

print(s)

for word in s:
    print(word)