#Case Study Title: "BookMart: A Mini Bookstore Management System"


from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main():
    while True:
        print("\n--- Welcome to BookMart! ---")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        
        choice = input("Enter your choice: ")
        
        if choice == '1':
            print("\n--- Book Management ---")
            print("1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            book_choice = input("Enter your choice: ")
            
            if book_choice == '1':
                try:
                    title = input("Title: ")
                    author = input("Author: ")
                    price = float(input("Price: "))
                    quantity = int(input("Quantity: "))
                    print(add_book(title, author, price, quantity))
                except ValueError as e:
                    print(f"Error: {e}")
            
            elif book_choice == '2':
                print("\nAvailable Books:")
                for book in view_books():
                    print(book)
            
            elif book_choice == '3':
                search_term = input("Enter book title or author to search: ")
                results = search_book(search_term)
                print("\nSearch Results:")
                for result in results:
                    print(result)

        elif choice == '2':
            print("\n--- Customer Management ---")
            print("1. Add Customer")
            print("2. View Customers")
            customer_choice = input("Enter your choice: ")
            
            if customer_choice == '1':
                try:
                    name = input("Name: ")
                    email = input("Email: ")
                    phone = input("Phone: ")
                    print(add_customer(name, email, phone))
                except ValueError as e:
                    print(f"Error: {e}")
            
            elif customer_choice == '2':
                print("\nCustomer List:")
                for customer in view_customers():
                    print(customer)

        elif choice == '3':
            print("\n--- Sales Management ---")
            print("1. Sell Book")
            print("2. View Sales Records")
            sales_choice = input("Enter your choice: ")
            
            if sales_choice == '1':
                try:
                    customer_name = input("Customer Name: ")
                    customer_email = input("Customer Email: ")
                    customer_phone = input("Customer Phone: ")
                    book_title = input("Book Title: ")
                    quantity = int(input("Quantity: "))
                    print(sell_book(customer_name, customer_email, customer_phone, book_title, quantity))
                except ValueError as e:
                    print(f"Error: {e}")
            
            elif sales_choice == '2':
                print("\nSales Records:")
                for record in view_sales():
                    print(record)

        elif choice == '4':
            print("Exiting the program. Thank you for using BookMart!")
            break
        
        else:
            print("Invalid choice! Please select a valid option.")

if __name__ == "__main__":
    main()
