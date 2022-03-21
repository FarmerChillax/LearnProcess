# -*- coding: utf-8 -*-
'''
    :file: main.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/03/21 10:53:03
'''

import re
import reprlib


class Sentence:
    def __init__(self, text) -> None:
        self.text = text
        self.words = re.compile('\w+').findall(text)

    def __getitem__(self, index):
        return self.words[index]

    def __len__(self):
        return len(self.words)

    def __repr__(self) -> str:
        return f'Sentence({reprlib.repr(self.text)})'

s = Sentence('"The time has come," the Walrus said,')
print(s)

for word in s:
    print(word)

print(list(s))