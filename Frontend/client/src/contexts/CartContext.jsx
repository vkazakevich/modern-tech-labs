import { createContext, useEffect, useState } from 'react'

const LOCAL_STORAGE_KEY = 'cart_product_ids'

export const CartContext = createContext([])

export const CartProvider = ({ children, ...props }) => {
  const [cartItems, setCartItems] = useState([])

  function getCartProductIds() {
    return JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY)) || []
  }

  function setCartProductIds(cartProductIds) {
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(cartProductIds))
  }

  useEffect(() => {
    setCartItems(getCartProductIds())
  }, [])

  const addProductToCart = async (productId) => {
    const items = [...cartItems, productId]
    setCartItems(items)
    setCartProductIds(items)
  }

  const cartItemsCount = cartItems.length

  const value = {
    cartItems,
    cartItemsCount,

    addProductToCart,
  }

  return <CartContext.Provider value={value}>{children}</CartContext.Provider>
}
