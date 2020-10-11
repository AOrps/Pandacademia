from flask import Flask, render_template, request, jsonify, url_for, redirect  #pip install flask
from back_end import *
app = Flask(__name__)

@app.route("/", methods = ['GET', 'POST'])
def index():
    if request.method == 'POST':
        nm = request.form.get('nm')
        feeling = request.form.get('feeling')
        temp = request.form.get('temp')
        sick = request.form.get('sick')
        sore = request.form.get('sore')
        contact = request.form.get('contact')
        location = request.form.get('location')
        return redirect(url_for('results', nm=nm, feeling=feeling, \
        temp=temp, sick=sick, sore=sore, contact=contact, \
        location=location))
    return render_template('index.html', title="🩺 Daily Screening")

@app.route("/info")
def info():
    return render_template("info.html", title="📚 CDC Info")

@app.route("/res")
def results(nm, feeling, temp, sick, sore, contact, location):
    ANSWER_SET = {}
    ANSWER_SET.add(feeling)
    ANSWER_SET.add(temp)
    ANSWER_SET.add(sick)
    ANSWER_SET.add(sore)
    ANSWER_SET.add(contact)
    ANSWER_SET.add(location)
    return render_template("results.html", code=num)

if __name__ == "__main__":
    app.run(port=8080)