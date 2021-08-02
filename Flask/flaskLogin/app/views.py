# -*- coding: utf-8 -*-
'''
    :file: views.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/01/31 15:36:48
'''
from flask import Blueprint, request
from flask_login import login_user, logout_user, current_user, login_required
from app.models import User

login_bp = Blueprint('login', __name__)

@login_bp.route("/login", methods=['GET', 'POST'])
def login():
    if current_user.is_authenticated:
        print(current_user.is_authenticated)
        return 'Was login'
    if request.method == 'GET':
        username = request.args.get('username')
        password = request.args.get('password')
        # return "GET method."
    else:
        username = request.form.get('username')
        password = request.form.get('password')
    
    # print(username, password)
    user = User.query.first()
    print(user.id, user.username)
    print(user.password_hash)
    if user:
        if username == user.username and user.validate_password(password):
            login_user(user, False)
            return 'login success'
        else:
            return 'No account.'
    return 'Not login.'

@login_bp.route('/logout')
@login_required
def logout():
    logout_user()
    return 'logout.'

@login_bp.route('/msg')
@login_required
def msg():
    
    return 'login msg.'