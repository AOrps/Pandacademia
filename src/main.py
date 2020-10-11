from flask import Flask, render_template  #pip install flask
from back_end import *
app = Flask(__name__)

@app.route("/")
def index():
    Proj_name = "Pandacademia.tech"
    return render_template("index.html", title="🩺 Daily Screening")


if __name__ == "__main__":
    app.run(port=8080)
