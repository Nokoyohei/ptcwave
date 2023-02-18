import mysql.connector
from card import Card
from datetime import datetime
import os
from dotenv import load_dotenv
load_dotenv()

conn = mysql.connector.connect(
  user='root',
  host='localhost',
  password=os.getenv('MYSQL_PASSWORD'),
  database='rush_price'
)

class CardRepository:
  def __init__(self):
    pass

  def create_cards(self, cards, date=None):
    fetch_date = date
    if fetch_date is None:
      fetch_date = datetime.now().strftime('%Y-%m-%d')
    sql = '''INSERT INTO cards 
      (
        name, 
        status, 
        rarity, 
        version, 
        type, 
        price, 
        fetch_date
      ) 
    VALUES 
      (%s, %s, %s, %s, %s, %s, %s)
    ON DUPLICATE KEY UPDATE price=VALUES(price)
    '''

    cards = [card.to_sql_object() + [fetch_date] for card in cards
      if card.to_sql_object() is not None and card.name != 'ウールー']

    with conn.cursor() as cursor:
      cursor.executemany(sql, cards)
      conn.commit()

