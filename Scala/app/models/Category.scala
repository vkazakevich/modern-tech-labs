package models

import play.api.libs.json._

case class Category (
    id: Option[Int], 
    name: String
)

object Category {
  implicit val categoryFormat: OFormat[Category] = Json.format[Category]
}