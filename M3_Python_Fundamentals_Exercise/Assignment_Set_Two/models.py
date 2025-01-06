from flask_sqlalchemy import SQLAlchemy
db = SQLAlchemy()

class Book(db.Model):
    __tablename__ = 'books'
    
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    title = db.Column(db.String(200), nullable=False)
    author = db.Column(db.String(200), nullable=False)
    published_year = db.Column(db.Integer, nullable=False)
    genre = db.Column(db.String(50), nullable=False)

    def __init__(self, title, author, published_year, genre):
        self.title = title
        self.author = author
        self.published_year = published_year
        self.genre = genre

    def __repr__(self):
        return f'<Book {self.title}>'
