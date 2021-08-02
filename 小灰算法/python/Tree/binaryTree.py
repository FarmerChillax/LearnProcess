# -*- coding: utf-8 -*-
'''
    :file: binaryTree.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/07/15 15:40:43
'''

class TreeNode:

    def __init__(self, data) -> None:
        self.data = data
        self.left = None
        self.right = None

def create_binary_tree(input_list:list=[]):
    """构建二叉树

    Args:
        input_list (list, optional): 输入数列. Defaults to [].
    """
    if input_list is None or len(input_list) == 0:
        return None
    
    data = input_list.pop(0)
    if data is None:
        return None
    
    node = TreeNode(data)
    node.left = create_binary_tree(input_list)
    node.right = create_binary_tree(input_list)
    return node

def pre_order_traversal(node:TreeNode):
    """前序遍历

    Args:
        node (TreeNode): 二叉树节点
    """
    if node is None:
        return
    
    print(node.data)
    pre_order_traversal(node.left)
    pre_order_traversal(node=node.right)
    return node

def in_order_traversal(node:TreeNode):
    """中序遍历

    Args:
        node (TreeNode): 二叉树节点
    """
    if node is None:
        return
    in_order_traversal(node.left)
    print(node.data)
    in_order_traversal(node.right)
    return node

def post_order_traversal(node:TreeNode):
    """后序遍历

    Args:
        node (TreeNode): 二叉树节点
    """
    if node is None:
        return
    post_order_traversal(node.left)
    post_order_traversal(node.right)
    print(node.data)
    return node

# with stack
def pre_order_traversal_with_stack(node:TreeNode):
    stack = []
    while node is not None or len(stack) > 0:
        print(node.data)
        stack.append(node)
        node = node.left
        if node == None and len(stack) > 0:
            node = stack.pop()
            node = node.right

if __name__ == '__main__':
    my_input_list = [3, 2, 9, None, None, 10, None, None, 8, None, 4]
    root = create_binary_tree(my_input_list)
    print(root.left.right.data)
    print('前序遍历')
    pre_order_traversal(root)
    # print('中序遍历')
    # in_order_traversal(root)
    # print('后序遍历')
    # post_order_traversal(root)
    print('stack前序遍历')
    pre_order_traversal(root)
