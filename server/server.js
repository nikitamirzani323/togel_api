const express = require('express')
const morgan = require('morgan')
const helmet = require('helmet')
const responseTime = require('response-time')
const createError = require('http-errors')
require('dotenv').config
const model = require('./Models/index')

const ApiController = require('./Routers/Apiroute')
const ServiceController = require('./Routers/Serviceroute')
const app = express()
const PORT = process.env.PORT || 9000
app.use(morgan('dev'))
app.use(helmet())
app.use(responseTime())
app.use(express.json())
app.use(express.urlencoded({
    extended: true
}));

app.use('/api', ApiController);
app.use('/service', ServiceController);

app.use(async (req,res,next) => {
    next(createError.NotFound('This route does not exist'))
});
app.use(async (err,req,res,next) => {
    res.status(err.status || 500)
    res.send({
        status: err.status || 500,
        message: err.message
    });
});




app.listen(PORT, () => {
    console.log(`Server running`);
});