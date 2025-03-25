package models

import play.api.libs.json.{Json, OFormat}
import scala.collection.mutable.ListBuffer

// (product_id, quantity)
case class Cart(id: Int, products: ListBuffer[(Int, Int)])

object Cart {
  implicit val format: OFormat[Cart] = Json.format[Cart]
}