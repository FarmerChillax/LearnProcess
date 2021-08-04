# -*- coding: utf-8 -*-
'''
    :file: views.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 22:19:15
'''

from flask.globals import g, current_app
from apiflask import input, APIBlueprint
from apiflask.decorators import auth_required, output
from app.database import users
from app.extension import auth
from app.schemas import UserRegInSchema, UserLoginSchema
from app.utils import generate_auth_token

bp = APIBlueprint("bp", __name__)

@bp.get("/")
@auth_required(auth)
def index():
    user = {"msg": "not found"}
    for item in users:
        if item["username"] == g.user:
            user = item
    return user

@bp.post("/reg")
@input(UserRegInSchema)
@output(UserRegInSchema(many=True))
def reg_user(data):
    """用户注册

    Args:
        data (dict): 注册信息
    """
    users.append(data)
    return users


@bp.post("/login")
@input(UserLoginSchema)
def login(data):
    token = None
    for user in users:
        if user["password"] == data["password"]:
            token = generate_auth_token(user)

    return {
        "token": token
    }

