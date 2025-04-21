import { useEffect, useState } from 'react'
import { Link } from 'react-router'
import { getProducts } from '../api/products'
import { useCart } from '../hooks/useCart'

function Products() {
  const [products, setProducts] = useState([])
  const { cartItemsCount, addProductToCart } = useCart()

  async function loadProducts() {
    const products = await getProducts()
    setProducts(products)
  }

  useEffect(() => {
    loadProducts()
  }, [])

  return (
    <>
      <h1>Products</h1>
      <div>
        Products in cart: {cartItemsCount}
        <br />
        <Link to="/cart">Cart</Link>
      </div>
      <ul>
        {products.map((product) => (
          <li key={product.ID}>
            {product.name}
            <br />
            Price: {product.price}
            <br />
            Quantity: {product.quantity}
            <br />
            <button onClick={() => addProductToCart(product.ID)}>
              Add to Cart
            </button>
            <br />
            <br />
          </li>
        ))}
      </ul>
    </>
  )
}

export default Products
