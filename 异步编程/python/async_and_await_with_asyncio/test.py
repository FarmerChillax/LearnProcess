# -*- coding: utf-8 -*-
'''
    :file: test.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/03/21 01:05:57
'''
import collections

Card = collections.namedtuple('Card', ['rank', 'suit'])

class FrenchDeck:
    """
    doctest模块测试
    >>> deck = FrenchDeck()
    >>> len(deck)
    52
    >>> for card in deck: # doctest: +ELLTPSIS
    ... print(card)
    Card(rank='2', suit='黑桃')
    ...
    """
    ranks = [str(n) for n in range(2, 11)] + list('JQKA')
    suits = '黑桃 红心 梅花 方块'.split()

    def __init__(self) -> None:
        self._cards = [Card(rank, suit) for suit in self.suits for rank in self.ranks]

    
    def __len__(self):
        return len(self._cards)

    def __getitem__(self, position):
        return self._cards[position]
    
    def __repr__(self) -> str:
        return f'FrenchDeck obj has {len(self._cards)}'

# from random import choice

deck = FrenchDeck()
print(deck)
# for card in deck:
#     print(card)

# if __name__ == '__main__':
#     import doctest
#     doctest.testmod()