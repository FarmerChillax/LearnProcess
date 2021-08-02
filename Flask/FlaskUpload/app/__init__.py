# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/02/12 17:27:09
'''
import os

from flask import Flask, json, jsonify, request
from flask_sqlalchemy import SQLAlchemy
# from flask_uploads import IMAGES, UploadSet, configure_uploads, patch_request_class
from werkzeug.utils import secure_filename

app = Flask(__name__)
app.config.from_pyfile('settings.py')

db = SQLAlchemy(app)

print(os.path.dirname(app.root_path))

app.config['UPLOAD_FOLDER'] = os.path.join(os.path.dirname(app.root_path), "fileUpload")
# UPLOAD_FOLDER
# files = UploadSet('file')
# configure_uploads(app, files)
# patch_request_class(app)

@app.route("/")
def index():
    filename = "../../../../test_.php"

    text = secure_filename(filename=filename)
    print(text)
    return jsonify({
        "point": secure_filename(request.args.get("value"))
    })

@app.route("/upload", methods=["POST"])
def upload():
    file = request.files.get('file')
    if file is None:
        return jsonify({'msg': 'No file part'})
    
    if file.filename == '':
        return jsonify({'msg': 'No selected file'})
    else:
        try:
            filename = secure_filename(file.filename)
            return jsonify({'msg': filename})
        except Exception as e:
            return jsonify({'msg': 'cathc error.'})
    return jsonify({
        "point": "index"
    })