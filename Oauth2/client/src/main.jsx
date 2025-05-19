import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router'

import { AuthProvider } from './providers/AuthProvider'

import Home from './pages/Home'
import Login from './pages/auth/Login'
import Register from './pages/auth/Register'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  </StrictMode>
)
