package models

import play.api.libs.json._

case class Product (
    id: Option[Int], 
    title: String, 
    description: Option[String], 
    price: BigDecimal
)

object Product {
  implicit val productFormat: OFormat[Product] = Json.format[Product]
}