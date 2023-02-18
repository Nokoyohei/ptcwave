import mysql.connector
from card import Card
from datetime import datetime

conn = mysql.connector.connect(
  user='root',
  host='localhost',
  database='rush_price'
)

class CardRepository:
  def __init__(self):
    pass

  def create_cards(self, cards, date=None):
    fetch_date = date
    if fetch_date is None:
      fetch_date = datetime.now().strftime('%Y-%m-%d')
    sql = '''INSERT INTO cards (
      name, 
      status, 
      rarity, 
      version, 
      type, 
      price, 
      fetch_date) VALUES (%s, %s, %s, %s, %s, %s, %s)
    '''

    cards = [card + [fetch_date] for card in cards]

    with conn.cursor() as cursor:
      cursor.executemany(sql, cards)
      conn.commit()

