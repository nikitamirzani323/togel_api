const sequilize = require('sequelize')
const db = new sequilize('db_togel','root','', {
    dialect: 'mysql',
    host: 'localhost'
})
module.exports = db