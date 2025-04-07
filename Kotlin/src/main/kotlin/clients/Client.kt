package clients

import models.Category
import models.categories
import models.products

abstract class Client {
    abstract suspend fun run()

    fun displayCategories() : String {
        println("Requested all categories")
        return categories.map { it.name }.joinToString(separator = "\n")
    }

    fun displayProductsByCategory(categoryName: String) : String {
        val category: Category? = categories.firstOrNull { it.name == categoryName }
        if (category == null) return "Incorrect category!"

        println("Requested products by category ${category.name}")
        return products
            .filter{ it.category == category }
            .map { "${it.name} (\$${it.price})" }
            .joinToString(separator = "\n")
    }
}