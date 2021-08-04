# -*- coding: utf-8 -*-
'''
    :file: gets.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 01:56:04
'''

from marshmallow.fields import Boolean, Integer, String
from sqlalchemy.sql.functions import user
from app.schemas import UserOutSchema
from app.model import User
from apiflask import APIBlueprint, abort
from apiflask.decorators import doc, output, input


user_bp = APIBlueprint("user", __name__)

@user_bp.get("/<int:id>")
@doc(summary="RESTAPI-通过id查询", description="用过id获取用户详细信息")
@output(UserOutSchema)
def get_user_info(id):
    """通过用户ID获取用户信息

    Args:
        id (int): 匹配传入URL中的id

    Returns:
        User: User模型
    """
    user = User.query.get(id)
    return user


@user_bp.get("/<string:name>")
@doc(summary="RESTAPI-通过name查询", description="通过用户名获取用户")
@output(UserOutSchema(many=True))
def get_users_info_byname(username):
    """通过用户名获取用户信息

    Args:
        username (string): e.g admin

    Returns:
        list: 查询结果
    """
    user = User.query.filter(User.name == username)
    return user

@user_bp.get("/all")
@doc(summary="RESTAPI-获取全部用户", description="通过用户名获取用户")
@output(UserOutSchema(many=True))
def get_all_user():
    users = User.query.all()
    return users

# get方法查询字符串

@user_bp.get("/query")
@doc(summary="查询字符串-查询用户", description="通过查询字符串的方式，选择通过 id 或 username 获取用户")
@input({"name": String(), "id": Integer()}, location='query')
@output(UserOutSchema(many=True))
def query_string_get_user(data):
    """通过查询字符串获取用户
         如果不传入参数，则获取全部用户数据

    Args:
        data (dict): 查询字符串内容

    Returns:
        user: user对象
    """
    id = data.get("id")
    name = data.get("name")
    if id:
        user = User.query.get(id)
    elif name:
        user = User.query.filter(User.name == name)
    else:
        user = User.query.all()

    return user

