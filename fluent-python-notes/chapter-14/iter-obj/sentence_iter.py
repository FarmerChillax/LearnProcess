# -*- coding: utf-8 -*-
'''
    :file: sentence_iter.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/06/08 12:34:20
'''

import re
import reprlib

RE_WORD = re.compile('\w+')

class Sentence:
    def __init__(self, text:str) -> None:
        self.text = text
        self.words = RE_WORD.findall(text)

    def __repr__(self) -> str:
        return f'Sentence({reprlib.repr(self.text)})' 

class SentenceIterator:
    def __init__(self, words) -> None:
        self.words = words
        self.index = 0
    
    def __next__(self):
        try:
            word = self.words[self.index]
        except IndexError:
            raise StopIteration()
        self.index += 1
        return word
    
    def __iter__(self):
        return self

print(Sentence("abc"))