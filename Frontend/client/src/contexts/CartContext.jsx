import PropTypes from 'prop-types'
import { createContext, useCallback, useEffect, useMemo, useState } from 'react'

const LOCAL_STORAGE_KEY = 'cart_product_ids'

export const CartContext = createContext([])

export const CartProvider = ({ children }) => {
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

  const addProductToCart = useCallback(
    async (productId) => {
      const items = [...cartItems, productId]
      setCartItems(items)
      setCartProductIds(items)
    },
    [cartItems]
  )

  const emptyCart = useCallback(() => {
    setCartItems([])
    setCartProductIds([])
  }, [])

  const cartItemsCount = cartItems.length

  const value = useMemo(
    () => ({
      cartItems,
      cartItemsCount,
      addProductToCart,
      emptyCart
    }),
    [cartItems, cartItemsCount, addProductToCart, emptyCart]
  )

  return <CartContext.Provider value={value}>{children}</CartContext.Provider>
}

CartProvider.propTypes = {
  children: PropTypes.node.isRequired
}
