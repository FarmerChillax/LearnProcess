from app import app, mail


@app.shell_context_processor
def make_shell_context():
    return dict(mail = mail)
