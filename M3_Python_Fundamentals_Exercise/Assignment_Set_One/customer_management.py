class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"

customers = [
    Customer("Remya", "remya@gmail.com", "9512763456"),
    Customer("Poorna", "poorna@gmail.com", "8447291569"),
    Customer("Swathee", "swathee@gmail.com", "9811357238"),
]

def add_customer(name, email, phone):
    if "@" not in email or "." not in email:
        raise ValueError("Invalid email address.")
    if not phone.isdigit() or len(phone) != 10:
        raise ValueError("Invalid phone number. Must be 10 digits.")
    new_customer = Customer(name, email, phone)
    customers.append(new_customer)
    return "Customer added successfully!"

def view_customers():
    return [customer.display_details() for customer in customers]
