const express = require('express')
const router = express.Router()
const ApiController = require('../Controllers/ApiController')

router.post('/serviceresult', ApiController.serviceresult)
router.post('/servicetoken', ApiController.servicetoken)
router.post('/serviceinit', ApiController.serviceinit)
router.post('/servicecheckpasaran', ApiController.servicecheckpasaran)

module.exports = router