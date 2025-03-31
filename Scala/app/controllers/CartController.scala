package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import play.api.libs.json._
import repositories.CartRepository
import models.Cart

import scala.concurrent.{ExecutionContext, Future}

@Singleton
class CartController @Inject() (
    val controllerComponents: ControllerComponents,
    cartRepository: CartRepository
)(implicit ec: ExecutionContext) extends BaseController {

  def showAll(): Action[AnyContent] = Action.async { implicit request =>
    cartRepository.list().map { cart_items =>
      Ok(Json.toJson(cart_items))
    }
  }

  def showById(id: Int): Action[AnyContent] = Action.async { implicit request =>
    cartRepository.findById(id).map {
      case Some(cart) => Ok(Json.toJson(cart))
      case None => NotFound(Json.obj("error" -> "Cart item not found"))
    }
  }

  def add(): Action[JsValue] = Action(parse.json).async { implicit request =>
    request.body
      .validate[Cart]
      .fold(
        errors => Future.successful(BadRequest(JsError.toJson(errors))),
        cart => {
          cartRepository.create(cart).map { created =>
            Created(Json.toJson(created))
          }
        }
      )
  }

  def update(id: Int): Action[JsValue] = Action(parse.json).async { implicit request =>
    request.body
      .validate[Cart]
      .fold(
        errors => Future.successful(BadRequest(JsError.toJson(errors))),
        cart => {
          cartRepository.update(id, cart).map {
            case 0 => NotFound(Json.obj("error" -> "Cart item not found"))
            case _ => Ok
          }
        }
      )
  }

  def delete(id: Int): Action[AnyContent] = Action.async {
    cartRepository.delete(id).map {
      case 0 => NotFound(Json.obj("error" -> "Cart item not found"))
      case _ => NoContent
    }
  }
}
