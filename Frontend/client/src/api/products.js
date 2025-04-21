export const getProducts = async function () {
  const url = 'http://localhost:8000/products/'

  try {
    const response = await fetch(url)
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`)
    }
    return await response.json()
  } catch (error) {
    console.error(error)
  }
}
