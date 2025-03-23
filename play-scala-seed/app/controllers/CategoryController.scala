package controllers

import javax.inject._
import play.api.libs.json._
import play.api.mvc._
import models.{DB, Category, CategoryData}

@Singleton
class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
    def getAll: Action[AnyContent] = Action {
      Ok(Json.toJson(DB.categories))
    }

    def getById(id: Int): Action[AnyContent] = Action {
      DB.categories.find(_.id == id) match {
        case Some(product) => Ok(Json.toJson(product))
        case None => NotFound(Json.obj("error" -> s"Category $id not found"))
      }
    }

    def add: Action[JsValue] = Action(parse.json) { request =>
      request.body.validate[CategoryData].fold(
        errors => BadRequest(Json.obj("error" -> "Invalid category format")),
        categoryData => {
          val newCategory = Category(DB.nextCategoryId, categoryData.name)
          DB.nextCategoryId += 1
          DB.categories += newCategory
          Created(Json.toJson(newCategory))
        }
      )
    }

    def update(id: Int): Action[JsValue] = Action(parse.json) { request =>
      request.body.validate[Category].fold(
        errors => BadRequest(Json.obj("error" -> "Invalid product format")),
        updatedCategory => {
          DB.categories.indexWhere(_.id == id) match {
            case -1 => NotFound(Json.obj("error" -> s"Category $id not found"))
            case idx =>
              DB.categories(idx) = updatedCategory
              Ok(Json.toJson(updatedCategory))
          }
        }
      )
    }

    def delete(id: Int): Action[AnyContent] = Action {
      DB.categories.indexWhere(_.id == id) match {
        case -1 => NotFound(Json.obj("error" -> s"Category $id not found"))
        case idx =>
          DB.categories.remove(idx)
          NoContent
      }
    }
}
