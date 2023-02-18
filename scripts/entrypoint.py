import csv
from card import Card
from card_repository import CardRepository
from fetch_rush_price import fetch_rush_price_from_spreadsheet
from datetime import datetime

now = datetime.now().strftime('%Y-%m-%d')
fetch_data_dir_path = 'fetch_data'

def get_unique_list(seq):
    seen = []
    return [x for x in seq if x not in seen and not seen.append(x)]

def write_csv(data, fetch_date):
  import csv

  with open(f'{fetch_data_dir_path}/{fetch_date}.csv', 'w', encoding='utf-8', newline='') as f:
    writer = csv.writer(f)
    writer.writerows(data)

def write_database_from_csv(date):
  cards = []
  with open(f'{fetch_data_dir_path}/{date}.csv') as f:
    reader = csv.reader(f)
    for row in reader:
      card = Card(
        name=row[0],
        status=row[1],
        rarity=row[2],
        version=row[3],
        type=row[4],
        price=row[5]
      )
      cards.append(card)

  cards = get_unique_list(cards)
  card_repository = CardRepository()
  card_repository.create_cards(cards, date)

d = fetch_rush_price_from_spreadsheet()
write_csv(d, now)
write_database_from_csv(now)