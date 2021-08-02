# -*- coding: utf-8 -*-
'''
    :file: binaryTreeAdjustment.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/07/31 09:01:49
'''

def up_adjust(array=[]):
    """二叉堆尾节点上浮

    Args:
        array (list, optional): 原数组. Defaults to [].
    """
    child_index = len(array) - 1
    parent_index = (child_index - 1) // 2
    # temp保存插入的叶子节点值，用于最后赋值
    temp = array[child_index]
    # 最小堆
    while child_index > 0 and temp < array[parent_index]:
        # 不用交换，单向赋值即可 (子节点上浮，父节点下沉)
        array[child_index] = array[parent_index]
        child_index = parent_index
        # 重新计算父节点的下标
        parent_index = (parent_index - 1) // 2
    
    array[child_index] = temp

def down_adjust(parent_index, length, array=[]):
    """二叉堆的节点下沉操作

    Args:
        parent_index (int): 下沉节点的下标
        length (int): 堆的长度范围
        array (list, optional): 原数组. Defaults to [].
    """
    # temp 保存父节点的值
    temp = array[parent_index]
    # 左孩子的下标
    child_index = 2 * parent_index + 1
    while child_index < length:
        # 如果有右孩子，且右孩子的值小于左孩子，则定位到右孩子
        if child_index + 1 < length and array[child_index+ 1] < array[child_index]:
            child_index += 1
        # 如果父节点的值小于任何一个孩子节点，直接退出
        if temp <= array[child_index]:
            break
        # 赋值操作
        array[parent_index] = array[child_index]
        parent_index = child_index
        child_index = 2 * parent_index + 1
    
    array[parent_index] = temp

def build_heap(array=[]):
    """二叉堆的构建

    Args:
        array (list, optional): 原数组. Defaults to [].
    """
    # 从最后一个非叶子节点开始，依次下沉
    for i in range((len(array) - 2) // 2, -1, -1):
        down_adjust(i, len(array), array=array)
    
my_array = list([1, 3, 2, 6, 5, 7, 8, 9, 10, 0])
up_adjust(my_array)
print(my_array)
my_array = [7, 1, 3, 10, 5, 2, 8, 9, 6]
build_heap(array=my_array)
print(my_array)