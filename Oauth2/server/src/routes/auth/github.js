import { Router } from 'express'
import { authUserByEmail, createUserIfNotExist } from '#services/auth'

const router = Router()

const GITHUB_OAUTH_URL = 'https://github.com/login/oauth/authorize'
const GITHUB_ACCESS_TOKEN_URL = 'https://github.com/login/oauth/access_token'

const GITHUB_USER_API = 'https://api.github.com/user'
const GITHUB_USER_EMAILS_API = 'https://api.github.com/user/emails'

const REDIRECT_URI = 'http://localhost:8000/auth/github/callback'

router.get('/auth/github', async (req, res) => {
  const params = new URLSearchParams({
    scope: 'user:email',
    client_id: process.env.GITHUB_CLIENT_ID,
    redirect_uri: REDIRECT_URI
  })

  res.redirect(`${GITHUB_OAUTH_URL}?${params}`)
})

router.get('/auth/github/callback', async (req, res) => {
  const { code } = req.query

  const response = await fetch(GITHUB_ACCESS_TOKEN_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json'
    },
    body: JSON.stringify({
      code,
      client_id: process.env.GITHUB_CLIENT_ID,
      client_secret: process.env.GITHUB_CLIENT_SECRET
    })
  })

  const { access_token } = await response.json()

  const userDataResponse = await fetch(GITHUB_USER_API, {
    headers: {
      Authorization: 'token ' + access_token
    }
  })
  const userData = await userDataResponse.json()

  let email = userData.email

  if (!email) {
    const userEmailsResponse = await fetch(GITHUB_USER_EMAILS_API, {
      headers: {
        Authorization: 'token ' + access_token
      }
    })

    const emails = await userEmailsResponse.json()

    if (!emails.length) {
      return res.status(500)
    }

    email = emails[0].email
  }

  await createUserIfNotExist({ fullName: userData.name, email: email })
  const token = await authUserByEmail(email)

  res.redirect(`http://localhost:5173/auth/token?token=${token}`)
})

export default router
