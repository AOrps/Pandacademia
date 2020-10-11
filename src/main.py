from flask import Flask, render_template  #pip install flask
from back_end import *
app = Flask(__name__)

@app.route("/")
def index():
    Proj_name = "Pandacademia.tech"
    return render_template("index.html", proj_name=Proj_name)



if __name__ == "__main__":
    app.run(port=8080)
