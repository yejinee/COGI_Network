const Squelize = require('sequelize');
const db = {}
const sequelize = new Squelize('hanium',
    'hanium',
    'mypassword',
{
    host : "haniumweb.cdxxehj8l1sk.ap-northeast-2.rds.amazonaws.com",
    dialect : 'mysql',
    operatorsAliases: false,

    pool : {
        max:5,
        min:0,
        acquire: 30000,
        idel : 10000
    }
}
    
)

db.sequelize = sequelize;
db.Squelize = Squelize;

module.exports = db;