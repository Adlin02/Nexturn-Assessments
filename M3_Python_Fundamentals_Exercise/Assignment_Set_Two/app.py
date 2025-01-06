from flask import Flask, request, jsonify, abort
from models import db, Book

app = Flask(__name__)

app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///books.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db.init_app(app)
with app.app_context():
    db.create_all()

@app.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    if not data or 'title' not in data or 'author' not in data or 'published_year' not in data or 'genre' not in data:
        abort(400, description="Missing required fields")

    new_book = Book(
        title=data['title'],
        author=data['author'],
        published_year=data['published_year'],
        genre=data['genre']
    )

    try:
        db.session.add(new_book)
        db.session.commit()
        return jsonify({
            'message': 'Book added successfully',
            'book_id': new_book.id
        }), 201
    except Exception as e:
        db.session.rollback()
        abort(500, description="Database error")

@app.route('/books', methods=['GET'])
def get_books():
    books = Book.query.all()
    if not books:
        return jsonify({"message": "No books available."}), 404
    return jsonify([{
        'id': book.id,
        'title': book.title,
        'author': book.author,
        'published_year': book.published_year,
        'genre': book.genre
    } for book in books])

@app.route('/books/<int:id>', methods=['GET'])
def get_book(id):
    book = Book.query.get(id)
    if not book:
        abort(404, description="No book exists with the provided ID")
    return jsonify({
        'id': book.id,
        'title': book.title,
        'author': book.author,
        'published_year': book.published_year,
        'genre': book.genre
    })

@app.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    book = Book.query.get(id)
    if not book:
        abort(404, description="No book exists with the provided ID")

    data = request.get_json()
    if 'title' in data:
        book.title = data['title']
    if 'author' in data:
        book.author = data['author']
    if 'published_year' in data:
        book.published_year = data['published_year']
    if 'genre' in data:
        book.genre = data['genre']

    try:
        db.session.commit()
        return jsonify({
            'message': 'Book updated successfully',
            'book_id': book.id
        })
    except Exception as e:
        db.session.rollback()
        abort(500, description="Database error")

@app.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    book = Book.query.get(id)
    if not book:
        abort(404, description="No book exists with the provided ID")

    try:
        db.session.delete(book)
        db.session.commit()
        return jsonify({'message': 'Book deleted successfully'}), 200
    except Exception as e:
        db.session.rollback()
        abort(500, description="Database error")

@app.errorhandler(404)
def not_found(error):
    return jsonify({
        'error': 'Book not found',
        'message': error.description
    }), 404

@app.errorhandler(400)
def bad_request(error):
    return jsonify({
        'error': 'Bad request',
        'message': error.description
    }), 400

@app.errorhandler(500)
def internal_error(error):
    return jsonify({
        'error': 'Internal server error',
        'message': error.description
    }), 500
if __name__ == "__main__":
    app.run(debug=True)
