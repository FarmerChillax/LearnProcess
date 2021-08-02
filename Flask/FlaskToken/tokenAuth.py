from flask import Flask, g
from flask_httpauth import HTTPTokenAuth

app = Flask(__name__)
auth = HTTPTokenAuth(scheme='Token')
tokens = {
    "xiaotao-1": "admin",
    "admin": "admin-value"
}

@auth.verify_token
def verify_token(token):
    if token in tokens:
        g.current_user = tokens[token]
        return True
    return False


@app.route('/')
@auth.login_required
def index():
    return "Hello, %s!" % g.current_user


if __name__ == '__main__':
    app.run()