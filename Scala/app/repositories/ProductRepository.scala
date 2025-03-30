package repositories

import javax.inject.{Inject, Singleton}
import play.api.db.slick.DatabaseConfigProvider
import slick.jdbc.JdbcProfile
import models.Product

import scala.concurrent.{ExecutionContext, Future}

@Singleton
class ProductRepository @Inject()(
    dbConfigProvider: DatabaseConfigProvider
)(implicit ec: ExecutionContext) {
  private val dbConfig = dbConfigProvider.get[JdbcProfile]
  
  import dbConfig._
  import profile.api._

  private class ProductTable(tag: Tag) extends Table[Product](tag, "products") {
    def id = column[Option[Long]]("id", O.PrimaryKey, O.AutoInc)
    def title = column[String]("title")
    def description = column[Option[String]]("description")
    def price = column[BigDecimal]("price")

    def * = (id, title, description, price) <> ((Product.apply _).tupled, Product.unapply)
  }

  private val products = TableQuery[ProductTable]

  def create(product: Product): Future[Product] = db.run {
    (products.map(p => (p.title, p.description, p.price))
        returning products.map(_.id)
        into { case ((title, description, price), id) =>
            Product(id, title, description, price)
        }
    ) += (product.title, product.description, product.price)
  }

  def list(): Future[Seq[Product]] = db.run {
    products.result
  }

  def findById(id: Long): Future[Option[Product]] = db.run {
    products.filter(_.id === id).result.headOption
  }
  def update(id: Long, product: Product): Future[Int] = db.run {
    products
      .filter(_.id === id)
      .map(p => (p.title, p.description, p.price))
      .update((product.title, product.description, product.price))
  }

  def delete(id: Long): Future[Int] = db.run {
    products.filter(_.id === id).delete
  }
}