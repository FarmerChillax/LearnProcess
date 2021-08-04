# -*- coding: utf-8 -*-
'''
    :file: app.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/04 22:48:55
'''
from apiflask import APIFlask, auth_required, input, output
from apiflask import HTTPTokenAuth
from flask.globals import g, current_app
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
from flask_cors import CORS
from schemas import UserRegInSchema, UserLoginSchema

users = [{
    "id": 1,
    "username": "farmer",
    "password": "admin",
    "url": "https://blog.farmer233.top"
}]

app = APIFlask(__name__)
app.config['SECRET_KEY'] = "YOUR_SECRET_KEY"

cors = CORS(app)
auth = HTTPTokenAuth()

@auth.verify_token
def verify_token(token):
    """验证token
        验证函数，验证token是否有效

    Args:
        token (str): 前端传来的token

    Returns:
        Bool: True表示token验证有效 否则False
    """
    g.user = None
    try:
        data = Serializer(current_app.config['SECRET_KEY']).loads(token)
    except Exception as e:
        return False
    if "username" in data:
        g.user = data["username"]
        return True
    return False

def generate_auth_token(payload:dict, expiration=3600 * 24 * 7):
    """生成token

    Args:
        payload (dict): token载体
        expiration (int, optional): token有效时间. Defaults to 3600*24*7.

    Returns:
        str: token令牌
    """
    s = Serializer(current_app.config["SECRET_KEY"], expires_in=expiration)
    return s.dumps(payload).decode()


@app.get("/")
@auth_required(auth)
def index():
    user = {"msg": "not found"}
    for item in users:
        if item["username"] == g.user:
            user = item
    return user

@app.post("/reg")
@input(UserRegInSchema)
@output(UserRegInSchema(many=True))
def reg_user(data):
    """用户注册

    Args:
        data (dict): 注册信息
    """
    users.append(data)
    return users


@app.post("/login")
@input(UserLoginSchema)
def login(data):
    token = None
    for user in users:
        if user["password"] == data["password"]:
            token = generate_auth_token(user)

    return {
        "token": token
    }


