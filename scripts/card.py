class Card:
  def __init__(self, name, status, rarity, version, type, price) -> None:
    self.name = name
    self.status = status
    self.rarity = rarity
    self.version = version
    self.type = type
    self.price = price

  def __str__(self) -> str:
    return f'name: {self.name}, status: {self.status}, rarity: {self.rarity}, version: {self.version}, type: {self.type}, price: {self.price}'

  def to_sql_object(self):
    if self.status == '' or self.status == '仕様':
      return None
    return [self.name, self.status, self.rarity, self.version, self.type, self.price]