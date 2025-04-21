import { useEffect, useState } from 'react'
import { getProducts } from '../api/products'
import { Link } from 'react-router'
import { useCart } from '../hooks/useCart'

function Cart() {
  const [products, setProducts] = useState([])
  const { cartItemsCount, cartItems } = useCart()

  async function loadProducts() {
    const products = await getProducts()
    setProducts(products)
  }

  useEffect(() => {
    loadProducts()
  }, [])

  function displayCartItem(productId) {
    const product = products.find((product) => product.ID === productId)
    if (!product) return <></>

    return (
      <li>
        Product: {product.name} | Price: ${product.price}
      </li>
    )
  }

  return (
    <>
      <h1>Cart ({cartItemsCount})</h1>
      {cartItemsCount > 0 ? (
        <div>
          <p>Products in cart:</p>
          <ol>{cartItems.map((productId) => displayCartItem(productId))}</ol>
          <p>
            Total: $
            {cartItems.reduce((acc, current) => {
              const product = products.find((p) => p.ID === current)
              return product ? (acc += product.price) : acc
            }, 0)}
          </p>
          <p>
            <Link to="/payments">Payments</Link>
          </p>
        </div>
      ) : (
        <div>Cart is empty</div>
      )}
    </>
  )
}

export default Cart
