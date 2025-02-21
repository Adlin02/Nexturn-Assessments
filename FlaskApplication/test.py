import pytest
from app import app, db, User

@pytest.fixture
def client():
    app.config['TESTING'] = True
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///:memory:'  # Fixed syntax error

    with app.test_client() as client:
        with app.app_context():
            db.create_all()
        yield client
        with app.app_context():
            db.drop_all()


def test_add_user(client):
    response = client.post('/users', json={'name': 'John Doe', 'email': 'john@example.com'})
    print(response.status_code, response.get_json())  # Debugging line

    assert response.status_code == 201
    assert response.get_json()['message'] == 'User added!'


def test_get_users(client):
    client.post('/users', json={'name': 'Jane Doe', 'email': 'jane@example.com'})
    response = client.get('/users')
    assert response.status_code == 200
    data = response.get_json()
    assert len(data) == 1
    assert data[0]['name'] == 'Jane Doe'


def test_get_user(client):
    client.post('/users', json={'name': 'Alice', 'email': 'alice@example.com'})
    response = client.get('/users/1')
    assert response.status_code == 200
    data = response.get_json()
    assert data['name'] == 'Alice'


def test_update_user(client):
    client.post('/users', json={'name': 'Bob', 'email': 'bob@example.com'})
    response = client.put('/users/1', json={'name': 'Bobby'})
    assert response.status_code == 200
    assert response.get_json()['message'] == 'User updated successfully!'

    # Verify update
    get_response = client.get('/users/1')
    assert get_response.get_json()['name'] == 'Bobby'


def test_delete_user(client):
    client.post('/users', json={'name': 'Charlie', 'email': 'charlie@example.com'})
    response = client.delete('/users/1')
    assert response.status_code == 200
    assert response.get_json()['message'] == 'User deleted!'

    # Verify deletion
    get_response = client.get('/users/1')
    assert get_response.status_code == 404


def test_404_error(client):
    response = client.get('/nonexistent')
    assert response.status_code == 404
    assert response.get_json()['error'] == 'Resource not found'


