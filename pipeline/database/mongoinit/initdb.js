db.products.drop();
db.products.insertMany([
    {
        "_id" : ObjectId("5fdcce9642085108206e5fed"),
        "name" : "Product Number One"
    },
    {
        "_id" : ObjectId("5fdd006e42085108206e5fee"),
        "name" : "Prodotto Numero Due"
    }
])
db.users.drop();
db.users.insertMany([
    {
        "_id": ObjectId("5fb7b1d521e4f91673dc2293"),
        "name": "First User",
        "username": "firstu",
        "password": "��I\u0000+Dܢ�r!:Z�\u000e0\""
    }
])