import { Router } from 'express'
import crypto from 'crypto'
import { authUserByEmail, createUserIfNotExist } from '#services/auth'

const router = Router()

const GOOGLE_OAUTH_URL = 'https://accounts.google.com/o/oauth2/v2/auth'
const GOOGLE_ACCESS_TOKEN_URL = 'https://oauth2.googleapis.com/token'
const GOOGLE_TOKEN_INFO_URL = 'https://oauth2.googleapis.com/tokeninfo'

const REDIRECT_URI = 'http://localhost:8000/auth/google/callback'

router.get('/auth/google', async (req, res) => {
  const scopes = [
    'https://www.googleapis.com/auth/userinfo.email',
    'https://www.googleapis.com/auth/userinfo.profile'
  ]

  const state = crypto.randomBytes(32).toString('hex')

  const params = new URLSearchParams({
    client_id: process.env.GOOGLE_CLIENT_ID,
    redirect_uri: REDIRECT_URI,
    response_type: 'code',
    scope: scopes.join(' '),
    access_type: 'offline',
    state: state
  })

  res.redirect(`${GOOGLE_OAUTH_URL}?${params}`)
})

router.get('/auth/google/callback', async (req, res) => {
  const { code } = req.query

  const params = {
    code,
    client_id: process.env.GOOGLE_CLIENT_ID,
    client_secret: process.env.GOOGLE_CLIENT_SECRET,
    redirect_uri: REDIRECT_URI,
    grant_type: 'authorization_code'
  }

  const response = await fetch(GOOGLE_ACCESS_TOKEN_URL, {
    method: 'POST',
    body: JSON.stringify(params)
  })

  const { id_token } = await response.json()
  const userInfoResponse = await fetch(
    `${GOOGLE_TOKEN_INFO_URL}?id_token=${id_token}`
  )
  const userInfo = await userInfoResponse.json()

  await createUserIfNotExist({ fullName: userInfo.name, email: userInfo.email })
  const token = await authUserByEmail(userInfo.email)

  res.redirect(`http://localhost:5173/auth/token?token=${token}`)
})

export default router
