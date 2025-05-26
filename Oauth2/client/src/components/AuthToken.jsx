import { Navigate, useNavigate } from 'react-router'
import { useAuth } from '../hooks/useAuth'
import { useEffect } from 'react'

function AuthToken() {
  const { auth } = useAuth()

  const navigate = useNavigate()

  useEffect(() => {
    const params = new URLSearchParams(window.location.search)
    const token = params.get('token')

    if (token) {
      auth({ token })
      navigate('/', { replace: true })
    } else {
      navigate('/login', { replace: true })
    }
  }, [auth, navigate])
}

export default AuthToken
