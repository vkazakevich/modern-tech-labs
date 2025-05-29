import prisma from '#libs/prisma'

export const GITHUB_PROVIDER = 'github'
export const GOOGLE_PROVIDER = 'google'

export const saveOauthCredentials = async (
  provider,
  userId,
  accessToken,
  expiresIn
) => {
  if (![GITHUB_PROVIDER, GOOGLE_PROVIDER].includes(provider)) {
    throw new Error('provider not exist')
  }

  await prisma.oauthCredentials.create({
    data: {
      provider,
      userId,
      accessToken,
      expireAt: new Date(Date.now() + expiresIn * 1000)
    }
  })
}

export const findOrCreateUser = async ({ fullName, email }) => {
  const user = await prisma.user.findUnique({
    where: { email: email }
  })

  if (user) return user

  const newUser = await prisma.user.create({
    data: {
      fullName,
      email,
      password: 'password'
    }
  })

  return newUser
}
