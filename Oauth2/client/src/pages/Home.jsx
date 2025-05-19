import { Navigate } from 'react-router'
import { useAuth } from '../hooks/useAuth'

function Home() {
  const { isAuth } = useAuth()

  if (!isAuth) {
    return <Navigate to="/login" replace />
  }

  return (
    <>
      <h1>Home</h1>
    </>
  )
}

export default Home
