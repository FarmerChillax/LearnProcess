from flask import Flask
from flask_httpauth import HTTPDigestAuth

app = Flask(__name__)

app.config['SECRET_KEY'] = 'Digest'
auth = HTTPDigestAuth()

users = {
    "admin": "admin",
    "xiaotao": "admin"
}

@auth.get_password
def get_password(username):
    if username in users:
        return users.get(username)
    return None

@app.route('/')
@auth.login_required
def index():
    print(auth.get_password)
    return 'Hello, %s!' % auth.username()

if __name__ == '__main__':
     app.run()