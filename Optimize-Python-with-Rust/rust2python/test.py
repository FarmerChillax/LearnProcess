# -*- coding: utf-8 -*-
'''
    :file: test.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/01/15 16:36:05
'''
import string_sum
import time

def target_func():
    def m(x):
        return x * x * x * x * x

    for a in range(0,10): 
        for b in range(0,10): 
            for c in range(0,10): 
                for d in range(0,10): 
                    for e in range(0,10): 
                        for f in range(0,10): 
                            if m(a) + m(b) + m(c) + m(d) + m(e) + m(f) == a * 10_0000 + b * 1_0000 + c * 1000 + d * 100 + e * 10 + f:
                                print(f"{a}{b}{c}{d}{e}{f}")



def test_fun(func, flage:str = "Python"):
    """计时函数

    Args:
        func (function): 回调函数

    """
    start = time.time()
    func()
    end = time.time()
    print(f"{flage} time: ",end - start)
    return start - end

# 开始计时 -> Rust
rust_time = test_fun(string_sum.test_func, flage="Rust")

# 测试纯Python
python_time = test_fun(target_func)

print(f"Rate: {python_time/rust_time}")
