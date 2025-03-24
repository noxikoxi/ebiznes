package models

import play.api.libs.json.{Json, OFormat}
import scala.collection.immutable.List

case class CartData(products: List[CartItemData])

object CartData {
  implicit val format: OFormat[CartData] = Json.format[CartData]
}
