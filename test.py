import os
import mysql.connector
from fastapi import FastAPI

app = FastAPI()

def _get_db_connection():
    return mysql.connector.connect(
        host="localhost",
        user="root",
        database="CarDatabase",
        password=os.environ.get("MYSQL_ROOT_PASSWORD"))


@app.get("/")
def read_root():
    return { "message": "Hello World!"}


@app.get("/cars")
def read_cars():
    conn = _get_db_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM Cars")
    cars = cursor.fetchall()
    cursor.close()
    conn.close()
    return cars