class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: ${self.price}, Quantity: {self.quantity}"

books = [
    Book("Verity", "Eiichiro Oda",13.99 , 10),
    Book("Silent Patient", "Masashi Kishimoto", 16.99, 5),
    Book("Secret", "Kubo",10.99, 8),
]

def add_book(title, author, price, quantity):
    if price <= 0 or quantity <= 0:
        raise ValueError("Price and quantity must be positive numbers.")
    new_book = Book(title, author, price, quantity)
    books.append(new_book)
    return "Book added successfully!"

def view_books():
    return [book.display_details() for book in books]

def search_book(search_term):
    found_books = [
        book.display_details() for book in books
        if search_term.lower() in book.title.lower() or search_term.lower() in book.author.lower()
    ]
    return found_books if found_books else ["No books found."]
