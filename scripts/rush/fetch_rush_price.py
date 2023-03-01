from requests_html import HTMLSession
session = HTMLSession()
headers = {
  "Cache-Control": "no-cache",
  "Pragma": "no-cache"
}

def parse_spreadsheet(url):
  resp = session.get(url, headers=headers)

  csv = []
  cnt = 0
  t = []
  for r in resp.html.find('td'):
    t.append(r.text)
    cnt += 1

    if cnt == 6:
      csv.append(t)
      t = []
      cnt = 0

  return csv

def fetch_rush_price_from_spreadsheet():
  standard = parse_spreadsheet('https://docs.google.com/spreadsheets/u/0/d/e/2PACX-1vQT3Q9qDbZUpnP3_WH2I5qw8O-U_PqXVhhoIzH2o-tSzeDND9FTuoGKbZiNHTbrzTgKAUA2_SvXFh_2/pubhtml/sheet?headers=false&gid=1490875147&range=A:F')
  old = parse_spreadsheet('https://docs.google.com/spreadsheets/u/0/d/e/2PACX-1vQT3Q9qDbZUpnP3_WH2I5qw8O-U_PqXVhhoIzH2o-tSzeDND9FTuoGKbZiNHTbrzTgKAUA2_SvXFh_2/pubhtml/sheet?headers=false&gid=159569114&range=A:F')

  return standard + old