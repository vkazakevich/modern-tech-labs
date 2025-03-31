package repositories

import javax.inject.{Inject, Singleton}
import play.api.db.slick.DatabaseConfigProvider
import slick.jdbc.JdbcProfile
import models.Cart

import scala.concurrent.{ExecutionContext, Future}

@Singleton
class CartRepository @Inject()(
    dbConfigProvider: DatabaseConfigProvider
)(implicit ec: ExecutionContext) {
  private val dbConfig = dbConfigProvider.get[JdbcProfile]
  
  import dbConfig._
  import profile.api._

  private class CartTable(tag: Tag) extends Table[Cart](tag, "cart") {
    def id = column[Option[Int]]("id", O.PrimaryKey, O.AutoInc)
    def product_id = column[Int]("product_id")
    def quantity = column[Int]("quantity")

    def * = (id, product_id, quantity) <> ((Cart.apply _).tupled, Cart.unapply)
  }

  private val cart_items = TableQuery[CartTable]

  def create(cart: Cart): Future[Cart] = db.run {
    (cart_items.map(p => (p.product_id, p.quantity))
        returning cart_items.map(_.id)
        into { case ((product_id, quantity), id) =>
            Cart(id, product_id, quantity)
        }
    ) += (cart.product_id, cart.quantity)
  }

  def list(): Future[Seq[Cart]] = db.run {
    cart_items.result
  }

  def findById(id: Int): Future[Option[Cart]] = db.run {
    cart_items.filter(_.id === id).result.headOption
  }
  def update(id: Int, cart: Cart): Future[Int] = db.run {
    cart_items
      .filter(_.id === id)
      .map(p => (p.product_id, p.quantity))
      .update((cart.product_id, cart.quantity))
  }

  def delete(id: Int): Future[Int] = db.run {
    cart_items.filter(_.id === id).delete
  }
}