# -*- coding: utf-8 -*-
'''
    :file: gets.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 01:56:04
'''

from app.schemas import UserOutSchema
from app.model import User
from apiflask import APIBlueprint
from apiflask.decorators import doc, output

user_bp = APIBlueprint("user", __name__)

@user_bp.get("/<int:id>")
@doc(summary="用户", description="用过id获取用户详细信息")
@output(UserOutSchema)
def get_user_info(id):
    user = User.query.get(id)
    return user


@user_bp.get("/all")
@doc(summary="用户", description="获取数据库所有用户")
@output(UserOutSchema(many=True))
def get_users_info():
    pass