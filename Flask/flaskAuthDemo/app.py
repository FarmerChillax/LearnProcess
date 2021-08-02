# -*- coding: utf-8 -*-
'''
    :file: app.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/09 18:09:03
'''
from flask import request, abort, jsonify, url_for, make_response
import click

from flask import Flask, g
from flask_httpauth import HTTPBasicAuth, HTTPTokenAuth, MultiAuth
from flask_sqlalchemy import SQLAlchemy

from werkzeug.security import generate_password_hash, check_password_hash
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
from itsdangerous import SignatureExpired, BadSignature

app = Flask(__name__)

app.config.from_pyfile('settings.py')

# 数据库配置
db = SQLAlchemy(app)

# commands


@app.cli.command()
@click.option('--drop', is_flag=True, help='Create after drop.')
def initdb(drop):
    '''Initialize the database.'''
    if drop:
        click.confirm(
            'This operation will delet the database, do you want to continue?', abort=True)
        db.drop_all()
        click.echo('Drop tables.')
    db.create_all()
    click.echo("init database.")


class User(db.Model):
    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(32), index=True)
    password_hash = db.Column(db.String(128))

    def hash_password(self, password):
        self.password_hash = generate_password_hash(password)

    def verify_password(self, password):
        return check_password_hash(self.password_hash, password)

    def generate_auth_token(self, expiration=600):
        s = Serializer('Bearer', expires_in=expiration)
        # print(s.dumps({'id': self.id}))
        return s.dumps({'id': self.id})

    @staticmethod
    def verify_auth_token(token):
        s = Serializer('Bearer')
        try:
            data = s.loads(token)
            print("data", data)
            print("token", token)
        except SignatureExpired:
            return None  # valid token, but expired
        except BadSignature:
            return None  # invalid token
        user = User.query.get(data['id'])
        return user


# 认证部分
basic_auth = HTTPBasicAuth()
token_auth = HTTPTokenAuth(scheme='Bearer')
multi_auth = MultiAuth(basic_auth, token_auth)
# @basic_auth.error_handler

datas = User.query.all()
# print(datas)  
for data in datas:
    s = Serializer(app.config['SECRET_KEY'])
    # token = s.dumps({'username': data['username']})
    print(data)
    token = data.generate_auth_token()
    print("User: " + data.username + "; password: " + data.password_hash)
    print("token:", token)

@basic_auth.verify_password
def verify_password(username, password):
    user = User.query.filter_by(username=username).first()
    if not user or not user.verify_password(password):
        return False
    g.user = user
    return True

@token_auth.error_handler
def unauthorized():
    return make_response(jsonify({'error': 'Unauthorized access'}), 401)

# Token Auth
@token_auth.verify_token
def verify_token(token):
    print(token)
    user = User.verify_auth_token(token=token)
    print(user)
    if user:
        g.user = user
        return True
    return False


# Token认证
@app.route('/api/token')
@multi_auth.login_required
def get_auth_token():
    token = g.user.generate_auth_token()
    print(token)
    return jsonify({ 'token': token.decode('ascii') })

# 普通认证views
@app.route('/api/users', methods=['POST'])
def new_user():
    username = request.json.get('username')
    password = request.json.get('password')
    if username is None or password is None:
        abort(400)  # missing arguments
    if User.query.filter_by(username=username).first() is not None:
        abort(400)  # existing user
    user = User(username=username)
    user.hash_password(password)
    db.session.add(user)
    db.session.commit()
    return jsonify({'username': user.username}), 201, {'Location': url_for('get_user', id=user.id, _external=True)}


@app.route('/api/resource')
@multi_auth.login_required
def get_resource():
    print(multi_auth.current_user())
    return jsonify({'data': 'Hello, %s!' % g.user.username})


@app.route('/api/get_user/<id>')
def get_user(id):
    return "get user views."


# "eyJhbGciOiJIUzUxMiIsImlhdCI6MTYxMjg5NjI3NSwiZXhwIjoxNjEyODk2ODc1fQ.eyJpZCI6Mn0.Ce-75PYHyFZ1pjfWhZ6y3o72gaBWZGD3cqelea-z6aoErM5a4vpxgz71Q83K0jj7VjXVf8qCs1u7Wj1_b2D52A"
# "eyJhbGciOiJIUzUxMiIsImlhdCI6MTYxMjg5NjI5NywiZXhwIjoxNjEyODk2ODk3fQ.eyJpZCI6Mn0.tKZ2j-1x0a6fyrQjl2zzADOpJavJ37eSipn5J1AnE1tOG9FsL5s1bVnc4LsBY-qD3uyIggvl8s6iDGJp6PGv5g"






# Base Auth
# @auth.verify_password
# def verify_password(username, password):
#     user = User.query.filter_by(username=username).first()
#     if not user or not user.verify_password(password):
#         return False
#     g.user = user
#     return True

# Base to Token Auth 
# @auth.verify_password
# def verify_password(username_or_token, password):
#     # first try to authenticate by token
#     user = User.verify_auth_token(username_or_token)
#     if not user:
#         # try to authenticate with username/password
#         user = User.query.filter_by(username = username_or_token).first()
#         if not user or not user.verify_password(password):
#             return False
#     g.user = user
#     return True
