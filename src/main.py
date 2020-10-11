from flask import Flask, render_template, request, jsonify, url_for, redirect  #pip install flask
from back_end import *
app = Flask(__name__)

@app.route("/", methods = ['GET', 'POST'])
def index():
    if request.method == 'POST':
        date = request.form.get('date')
        return redirect(url_for('booking', date=date))
    return render_template('index.html', title="🩺 Daily Screening")

@app.route("/info")
def info():
    return render_template("info.html", title="📚 CDC Info")

@app.route("/res")
def results():
    num=1
    return render_template("results.html", proj_name=PROJ_NAME, code=num)

if __name__ == "__main__":
    app.run(port=8080)