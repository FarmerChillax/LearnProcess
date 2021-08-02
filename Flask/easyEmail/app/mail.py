from threading import Thread
from flask_mail import Message
from app import mail, app

def _send_async_mail(app, message):
    with app.app_context():
        mail.send(message)

def send_mail(subject, to, body) -> Thread:
    message = Message(subject=subject, recipients=[to], body=body)
    thr = Thread(target=_send_async_mail, args=[app, message])
    thr.start()
    return thr