from book_management import books
from customer_management import customers

sales_records = []

def sell_book(customer_name, customer_email, customer_phone, book_title, quantity):
    for book in books:
        if book.title.lower() == book_title.lower():
            if book.quantity >= quantity:
                book.quantity -= quantity
                sales_records.append({
                    "customer_name": customer_name,
                    "customer_email": customer_email,
                    "customer_phone": customer_phone,
                    "book_title": book.title,
                    "quantity_sold": quantity,
                })
                return f"Sale successful! Remaining quantity: {book.quantity}"
            else:
                return f"Error: Only {book.quantity} copies available. Sale cannot be completed."
    return "Error: Book not found."

def view_sales():
    return [
        f"Customer: {record['customer_name']}, Book Title: {record['book_title']}, Quantity Sold: {record['quantity_sold']}"
        for record in sales_records
    ]
