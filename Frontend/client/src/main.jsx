import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router'
import './index.css'
import Products from './components/Products.jsx'
import Payments from './components/Payments.jsx'
import Cart from './components/Cart.jsx'
import { CartProvider } from './contexts/CartContext.jsx'

import './plugins'

createRoot(document.getElementById('root')).render(
  <CartProvider>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Products />} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/payments" element={<Payments />} />
      </Routes>
    </BrowserRouter>
  </CartProvider>
)
