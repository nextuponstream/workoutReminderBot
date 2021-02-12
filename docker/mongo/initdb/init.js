db.createUser({
    user: 'root', // FIXME use .env
    pwd: 'bwr123qwe', // FIXME use .env
    roles: [
        {
            role: 'readWrite',
            db: 'my_mongo_db', // FIXME use .env
        },
    ],
});

db = new Mongo().getDB("my_mongo_db"); // FIXME use .env

db.createCollection('activities', { capped: false });