import { Navigate } from 'react-router'
import { useAuth } from '../hooks/useAuth'

function Home() {
  const { isAuth, logout } = useAuth()

  if (!isAuth) {
    return <Navigate to="/login" replace />
  }

  return (
    <>
      <p className="p-3">You're authorized!</p>
      <p className="p-3">
        <button
          onClick={logout}
          className="border rounded-xl border-red-500 p-3"
        >
          Logout
        </button>
      </p>
    </>
  )
}

export default Home
