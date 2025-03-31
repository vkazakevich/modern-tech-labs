package controllers

import javax.inject._
import play.api._
import play.api.mvc._
import play.api.libs.json._
import repositories.CategoryRepository
import models.Category

import scala.concurrent.{ExecutionContext, Future}

@Singleton
class CategoriesController @Inject() (
    val controllerComponents: ControllerComponents,
    categoryRepository: CategoryRepository
)(implicit ec: ExecutionContext) extends BaseController {

  def showAll(): Action[AnyContent] = Action.async { implicit request =>
    categoryRepository.list().map { categories =>
      Ok(Json.toJson(categories))
    }
  }

  def showById(id: Int): Action[AnyContent] = Action.async { implicit request =>
    categoryRepository.findById(id).map {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(Json.obj("error" -> "Category not found"))
    }
  }

  def add(): Action[JsValue] = Action(parse.json).async { implicit request =>
    request.body
      .validate[Category]
      .fold(
        errors => Future.successful(BadRequest(JsError.toJson(errors))),
        category => {
          categoryRepository.create(category).map { created =>
            Created(Json.toJson(created))
          }
        }
      )
  }

  def update(id: Int): Action[JsValue] = Action(parse.json).async { implicit request =>
    request.body
      .validate[Category]
      .fold(
        errors => Future.successful(BadRequest(JsError.toJson(errors))),
        category => {
          categoryRepository.update(id, category).map {
            case 0 => NotFound(Json.obj("error" -> "Category not found"))
            case _ => Ok
          }
        }
      )
  }

  def delete(id: Int): Action[AnyContent] = Action.async {
    categoryRepository.delete(id).map {
      case 0 => NotFound(Json.obj("error" -> "Category not found"))
      case _ => NoContent
    }
  }
}
