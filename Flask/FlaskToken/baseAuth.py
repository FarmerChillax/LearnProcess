from flask import Flask
from flask_httpauth import HTTPBasicAuth
from werkzeug.security import generate_password_hash, check_password_hash

app = Flask(__name__)
auth = HTTPBasicAuth()

users = {
    "admin": generate_password_hash("admin"),
    "susan": generate_password_hash("bye")
}

@auth.verify_password
def verify_password(username, password):
    if username in users and check_password_hash(users.get(username), password):
        return username

@auth.hash_password
def hash_pw(password):
    return generate_password_hash(password)


@app.route('/')
@auth.login_required
def index():
    print(auth.hash_password_callback)
    print(auth.verify_password_callback)
    print(users.get('admin'))
    print("Hello, %s!" % auth.username())
    print("password:", )
    return "Hello, {}!".format(auth.current_user())


if __name__ == '__main__':
    app.run()