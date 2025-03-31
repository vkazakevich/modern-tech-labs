package models

import play.api.libs.json._

case class Cart (
    id: Option[Int], 
    product_id: Int,
    quantity: Int
)

object Cart {
  implicit val cartFormat: OFormat[Cart] = Json.format[Cart]
}