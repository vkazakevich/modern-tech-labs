package repositories

import javax.inject.{Inject, Singleton}
import play.api.db.slick.DatabaseConfigProvider
import slick.jdbc.JdbcProfile
import models.Category

import scala.concurrent.{ExecutionContext, Future}

@Singleton
class CategoryRepository @Inject()(
    dbConfigProvider: DatabaseConfigProvider
)(implicit ec: ExecutionContext) {
  private val dbConfig = dbConfigProvider.get[JdbcProfile]
  
  import dbConfig._
  import profile.api._

  private class CategoryTable(tag: Tag) extends Table[Category](tag, "categories") {
    def id = column[Option[Int]]("id", O.PrimaryKey, O.AutoInc)
    def name = column[String]("name")

    def * = (id, name) <> ((Category.apply _).tupled, Category.unapply)
  }

  private val categories = TableQuery[CategoryTable]

  def create(category: Category): Future[Category] = db.run {
    (categories.map(p => p.name)
        returning categories.map(_.id)
        into { case ((name), id) =>
            Category(id, name)
        }
    ) += (category.name)
  }

  def list(): Future[Seq[Category]] = db.run {
    categories.result
  }

  def findById(id: Int): Future[Option[Category]] = db.run {
    categories.filter(_.id === id).result.headOption
  }
  def update(id: Int, category: Category): Future[Int] = db.run {
    categories
      .filter(_.id === id)
      .map(p => (p.name))
      .update((category.name))
  }

  def delete(id: Int): Future[Int] = db.run {
    categories.filter(_.id === id).delete
  }
}