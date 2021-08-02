# -*- coding: utf-8 -*-
'''
    :file: users.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/06/04 15:20:30
'''
from flask import Blueprint, json, jsonify, render_template
from flask.globals import current_app, request, g
from flask.helpers import url_for
from flask_sqlalchemy import Pagination
from itsdangerous import exc
from sqlalchemy.sql.functions import user
from app.models import User
from app.extensions import auth


user_bp = Blueprint('user', __name__)

@user_bp.route("/")
def index():
    return "User Index"

@user_bp.route("/print")
def myprint():
    a = url_for('user.myprint')
    print(a)
    return url_for('user.index')


@user_bp.route("/<username>")
def get_user(username):
    print(username)
    user:User = User.query.filter(User.username == username).first()
    print(user)
    if user != None:
        return jsonify({
            "username": user.username
        })
    return jsonify({
        "username": "404"
    })


@user_bp.route('/login', methods=['GET', 'POST'])
def login():
    try:
        data = request.get_json()
        username = data.get("username", None)
        password = data.get('password', None)
        if username is None or password is None:
            return jsonify({
                "msg": "参数不全"
            })
        user:User = User.query.filter(User.username == username).first()
        if user.validate_password(password):
            token = user.generate_auth_token()
            return jsonify({
                "msg": token
            })
        return jsonify({
                "msg": "密码错误"
            })
    except Exception as e:
        current_app.logger.error(e)

@user_bp.route("/check")
@auth.login_required
def check():
    print(g.username)
    return "None"

@user_bp.route("/args")
def args():
    per_page = 10
    page = request.args.get('page', 1)
    pagination:Pagination = User.query.paginate(page, per_page=per_page)
    print(pagination)
    users = pagination.items
    print(users)
    return render_template('index.html', pagination=pagination, args={
        "brank": [123, 233]
    })
    # res = url_for('user.args', )
    # return res