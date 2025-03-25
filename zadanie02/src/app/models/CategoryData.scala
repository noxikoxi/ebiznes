package models

import play.api.libs.json.{Json, OFormat}

case class CategoryData(name: String)

object CategoryData {
  implicit val format: OFormat[CategoryData] = Json.format[CategoryData]
}
