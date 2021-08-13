const sequilize = require('sequelize')
const db = new sequilize('db_tot','sperma','asdQWE123!@#', {
    dialect: 'mysql',
    host: '165.22.242.64'
})
module.exports = db