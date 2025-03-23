package models

import play.api.libs.json.{Json, OFormat}
import scala.collection.mutable.ListBuffer

case class CartItemData(product_id: Int, quantity: Int)

object CartItemData {
  implicit val format: OFormat[CartItemData] = Json.format[CartItemData]
}
