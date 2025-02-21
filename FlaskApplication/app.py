from flask import Flask, request, jsonify, render_template
from flask_sqlalchemy import SQLAlchemy
from flask_jwt_extended import JWTManager
from flask_httpauth import HTTPBasicAuth

app = Flask(__name__)
auth = HTTPBasicAuth()

# Simple user data for authentication
users = {
    "admin": "password123",
    "user": "mypassword"
}

# Authentication handler
@auth.verify_password
def verify_password(username, password):
    if username in users and users[username] == password:
        return username
    return None

# Configurations
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///users.db'
app.config['SECRET_API_KEY'] = '858b3b10d6539b1ac515146a9b57fc8f'


# Initialize Extensions
db = SQLAlchemy(app)
jwt = JWTManager(app)

@app.route('/')
def home():
    return render_template('index.html')

# Define User model
class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(50), nullable=False)
    email = db.Column(db.String(50), unique=True, nullable=False)

# 🔐 Auth-protected Route
@app.route('/auth/users')
@auth.login_required
def get_protected_users():
    return jsonify({"message": f"Hello, {auth.current_user()}! Welcome to the protected route."})

# ➕ Add User
@app.route('/users', methods=['POST'])
def add_user():
    data = request.get_json()
    new_user = User(name=data['name'], email=data['email'])
    db.session.add(new_user)
    db.session.commit()
    return jsonify({"message": "User added!"}), 201

# 📋 Get All Users
@app.route('/users', methods=['GET'])
def get_users():
    users = User.query.all()
    output = [{"id": user.id, "name": user.name, "email": user.email} for user in users]
    return jsonify(output)

# 🔍 Get Single User
@app.route('/users/<int:id>', methods=['GET'])
def get_user(id):
    user = User.query.get_or_404(id)
    return jsonify({"id": user.id, "name": user.name, "email": user.email})

# ✏️ Update User
@app.route('/users/<int:id>', methods=['PUT'])
def update_user(id):
    user = User.query.get_or_404(id)
    data = request.get_json()

    user.name = data.get('name', user.name)
    user.email = data.get('email', user.email)

    db.session.commit()
    return jsonify({"message": "User updated successfully!"})

# ❌ Delete User
@app.route('/users/<int:id>', methods=['DELETE'])
def delete_user(id):
    user = User.query.get_or_404(id)
    db.session.delete(user)
    db.session.commit()
    return jsonify({"message": "User deleted!"})

# 🚨 Error Handling
@app.errorhandler(404)
def not_found(error):
    return jsonify({"error": "Resource not found"}), 404

@app.errorhandler(500)
def server_error(error):
    return jsonify({"error": "Internal server error"}), 500

if __name__ == '__main__':
    with app.app_context():
        db.create_all()
    app.run(host='0.0.0.0', port=5000, debug=True)
