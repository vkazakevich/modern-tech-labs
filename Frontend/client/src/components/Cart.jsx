import { useEffect, useState } from 'react'
import { getProducts } from '../api/products'
import { Link } from 'react-router'

function Cart() {
  const [products, setProducts] = useState([])

  function getCartProductIds() {
    return JSON.parse(localStorage.getItem('cart_product_ids')) || []
  }

  async function loadProducts() {
    const cartProductIds = getCartProductIds()
    const products = await getProducts()

    setProducts(
      products.filter((product) => cartProductIds.includes(product.ID))
    )
  }

  useEffect(() => {
    loadProducts()
  }, [])

  return (
    <>
      <h1>Cart</h1>
      {products.length > 0 ? (
        <div>
          <p>Products in cart:</p>
          <ul>
            {products.map((product) => (
              <li key={product.ID}>
                Product: {product.name} | Price: ${product.price}
              </li>
            ))}
          </ul>
          <p>
            Total: $
            {products.reduce((acc, current) => {
              return (acc += current.price)
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
