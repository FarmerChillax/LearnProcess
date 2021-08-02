from app import db

class Mail(db.Model):
    id = db.Column(db.Integer,  primary_key=True)
    recipient = db.Column(db.String(50), index = True)
    recipient_email = db.Column(db.String(50))
    subject = db.Column(db.String(50))
    body = db.Column(db.Text)

def write_db_email(req_data):
    """
    电子邮件存档
    """
    mail = Mail()
    mail.recipient = req_data.get('recipient')
    mail.recipient_email = req_data.get('user_email')
    mail.subject = req_data.get("subject")
    mail.body = req_data.get('body')
    db.session.add(mail)
    db.session.commit()