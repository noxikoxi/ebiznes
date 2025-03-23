package models

import play.api.libs.json.{Json, OFormat}
import scala.collection.mutable.ListBuffer

case class CartData(products: ListBuffer[(Int, Int)])

object CartData {
  implicit val format: OFormat[CartData] = Json.format[CartData]
}
