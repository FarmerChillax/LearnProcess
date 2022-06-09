# -*- coding: utf-8 -*-
'''
    :file: sentence_gen.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/06/08 12:54:13
'''
import re
import reprlib

RE_WORD = re.compile("\w+")

class Sentence:
    def __init__(self, text:str) -> None:
        self.text = text
        self.words = RE_WORD.findall(text)

    def __repr__(self) -> str:
        return f'Sentence({reprlib.repr(self.text)})' 
    
    def __iter__(self):
        for word in self.words:
            yield word
        return

