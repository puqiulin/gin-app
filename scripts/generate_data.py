import random

import records
from faker import Faker
from urllib.parse import quote_plus
import psycopg2
from psycopg2 import connect, OperationalError

fake = Faker()


# password = quote_plus("@123456")
# db = records.Database(f"postgresql+psycopg2://postgres:{password}@127.0.0.1:5432/gin-app")
#
# def generate_users(num):
#     users = [{"name": fake.name(), "email": fake.email()} for _ in range(num)]
#     query = 'insert into users (name, email) values (:name, :email)'
#     db.bulk_query(query, users)
#
# generate_users(100)

def create_db():
    try:
        conn = connect(
            dbname="gin-app",
            user="postgres",
            password="@123456",
            host="127.0.0.1",
            port="5432"
        )
        return conn
    except OperationalError as e:
        print(f"An error occurred: {e}")
        return None


db = create_db()


def generate_users(num):
    users = [(fake.name(), fake.email(), fake.phone_number(), random.choice([0, 1])) for _ in range(num)]
    query = 'insert into users (name, email,phone,gender) values (%s, %s,%s,%s)'
    cursor = db.cursor()

    try:
        cursor.executemany(query, users)
        db.commit()
        print(f"{cursor.rowcount} records inserted successfully.")
    except (Exception, psycopg2.DatabaseError) as error:
        print(f"Error: {error}")
        db.rollback()
    finally:
        cursor.close()


generate_users(100)
