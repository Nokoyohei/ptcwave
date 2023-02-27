from requests_html import HTMLSession
import re
from operator import itemgetter 

session = HTMLSession()
headers = {
  "Cache-Control": "no-cache",
  "Pragma": "no-cache"
}

def parse_table(url):
  resp = session.get(url, headers=headers)

  contents = resp.html.find('.table_main')
  csv = []
  for r in contents:
    card_info = r.find('.table_left_cell', first=True).text
    price = r.find('.table_right_cell', first=True).text

    csv_base_card_info = re.findall(r'^(.*?)[(](.*?)[)][{](.*?)[}][〈](.*?)[〉][\[](.*?)[\]]', card_info)
    if len(csv_base_card_info) == 0:
      csv_base_card_info = [[card_info, '-', '-', '-', '-']]

    csv.append(list(csv_base_card_info[0]) + [price])

  return csv

def fetch_hareruya_price_from_site():
  ss = parse_table('https://www.hareruya2.com/page/17')
  sm = parse_table('https://www.hareruya2.com/page/18')
  xy = parse_table('https://www.hareruya2.com/page/19')
  bw = parse_table('https://www.hareruya2.com/page/20')
  sv = parse_table('https://www.hareruya2.com/page/55')

  return ss + sm + xy + bw + sv