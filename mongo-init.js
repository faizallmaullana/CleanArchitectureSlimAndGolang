db = db.getSiblingDB('admin');

db.createUser({
  user: 'app_user',
  pwd: 'app_password',
  roles: [
    { role: 'readWrite', db: 'app_db' }
  ]
});