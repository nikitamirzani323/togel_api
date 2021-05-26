const sequilize = require('sequelize')
const db = new sequilize('db_togel2','root','', {
    dialect: 'mysql',
    host: 'localhost'
})
module.exports = db