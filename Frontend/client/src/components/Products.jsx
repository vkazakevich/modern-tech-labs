import { useEffect, useState } from 'react'

function Products() {
  const [products, setProducts] = useState([])

  async function getProducts() {
    const url = 'http://localhost:8000/products/'
    try {
      const response = await fetch(url)
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`)
      }
      const products = await response.json()
      setProducts(products)
    } catch (error) {
      console.error(error)
    }
  }

  useEffect(() => {
    getProducts()
  }, [])

  return (
    <>
      <h1>Products</h1>
      <ul>
        {products.map((product) => (
          <li>
            {product.name}
            <br />
            Price: {product.price}
            <br />
            Quantity: {product.quantity}
          </li>
        ))}
      </ul>
    </>
  )
}

export default Products
