import { Router } from 'express'

import apiAuth from '#routes/api/auth'

import authGoogle from '#routes/auth/google'
import authGithub from '#routes/auth/github'

const router = Router()

router.use('/api', apiAuth)

router.use('/', authGoogle)
router.use('/', authGithub)

export default router
