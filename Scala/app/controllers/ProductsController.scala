package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import play.api.libs.json._
import repositories.ProductRepository
import models.Product

import scala.concurrent.{ExecutionContext, Future}

@Singleton
class ProductsController @Inject() (
    val controllerComponents: ControllerComponents,
    productRepository: ProductRepository
)(implicit ec: ExecutionContext) extends BaseController {

  def showAll(): Action[AnyContent] = Action.async { implicit request =>
    productRepository.list().map { products: Seq[Product] =>
      Ok(Json.toJson(products))
    }
  }

  def showById(id: Long): Action[AnyContent] = Action.async { implicit request =>
    productRepository.findById(id).map {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(Json.obj("error" -> "Product not found"))
    }
  }

  def add(): Action[JsValue] = Action(parse.json).async { implicit request =>
    request.body
      .validate[Product]
      .fold(
        errors => Future.successful(BadRequest(JsError.toJson(errors))),
        product => {
          productRepository.create(product).map { created =>
            Created(Json.toJson(created))
          }
        }
      )
  }

  def update(id: Long): Action[JsValue] = Action(parse.json).async { implicit request =>
    request.body
      .validate[Product]
      .fold(
        errors => Future.successful(BadRequest(JsError.toJson(errors))),
        product => {
          productRepository.update(id, product).map {
            case 0 => NotFound(Json.obj("error" -> "Product not found"))
            case _ => Ok
          }
        }
      )
  }

  def delete(id: Long): Action[AnyContent] = Action.async {
    productRepository.delete(id).map {
      case 0 => NotFound(Json.obj("error" -> "Product not found"))
      case _ => NoContent
    }
  }
}
