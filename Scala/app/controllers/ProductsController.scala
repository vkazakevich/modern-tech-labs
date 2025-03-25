package controllers

import javax.inject._
import play.api._
import play.api.mvc._

@Singleton
class ProductsController @Inject() (
    val controllerComponents: ControllerComponents
) extends BaseController {
  def index() = Action { implicit request: Request[AnyContent] =>
    Ok("Products")
  }

  def show(id: Long) = Action { implicit request: Request[AnyContent] =>
    Ok("Product " + id)
  }
}
