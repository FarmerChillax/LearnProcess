from flask import request, jsonify
from flask_mail import Message

from app import app, mail, db
from app.models import Mail

@app.route("/mail")
def send_mail():
    req_data = request.json
    
    if req_data.get("user") == None:
        message = Message(subject=req_data.get('subject'), recipients=[req_data.get('user_email')], body=req_data.get('body'))
    else:
        message = Message(subject=req_data.get('subject'), recipients=[req_data.get('user') + " " + req_data.get('user_email')], body=req_data.get('body'))
    mail.send(message)

    db_mail = Mail()
    db_mail.recipient = req_data.get("user")
    db_mail.recipient_email = req_data.get('user_email')
    db_mail.subject = req_data.get('subject')
    db_mail.body = req_data.get('body')
    db.session.add(db_mail)
    db.session.commit()

    res_data = req_data
    return jsonify(res_data)