package models

import play.api.libs.json.{Json, OFormat}

case class ProductData(name: String, price: Double, category_id: Int)

object ProductData {
  implicit val format: OFormat[ProductData] = Json.format[ProductData]
}