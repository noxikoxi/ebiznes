package controllers

import javax.inject.*
import play.api.libs.json.*
import play.api.mvc.*
import models.{DB, Product, ProductData}

import scala.collection.mutable.ListBuffer

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

    def getAll: Action[AnyContent] = Action {
      val productsWithCategory = DB.products.map { product =>
        val categoryName = DB.categories.find(_.id == product.category_id).map(_.name).getOrElse("Unknown")

        Json.obj(
          "id" -> product.id,
          "name" -> product.name,
          "price" -> product.price,
          "category" -> categoryName
        )
      }
      Ok(Json.toJson(productsWithCategory))
    }

    def getById(id: Int): Action[AnyContent] = Action {
      DB.products.find(_.id == id) match {
        case Some(product) => Ok(Json.obj(
          "id" -> product.id,
          "name" -> product.name,
          "price" -> product.price,
          "category" -> DB.categories.find(_.id == product.category_id).map(_.name).getOrElse("Unknown")
        ))
        case None => NotFound(Json.obj("error" -> s"Product $id not found"))
      }
    }

    def add: Action[JsValue] = Action(parse.json) { request =>
      request.body.validate[ProductData].fold(
        errors => BadRequest(Json.obj("error" -> "Invalid product format")),
        productData => {
          val newProduct = Product(DB.nextProductId, productData.name, productData.price, productData.category_id)
          DB.nextProductId += 1
          DB.products += newProduct
          Created(Json.toJson(newProduct))
        }
      )
    }

    def update(id: Int): Action[JsValue] = Action(parse.json) { request =>

      DB.products.indexWhere(_.id == id) match {
        case -1 => NotFound(Json.obj("error" -> s"Product $id not found"))
        case idx =>
          request.body.validate[ProductData].fold(
            errors => BadRequest(Json.obj("error" -> "Invalid product format")),
            productData => {
                val updatedProduct = Product(id, productData.name, productData.price, productData.category_id)
                DB.products(idx) = updatedProduct
                Ok(Json.toJson(updatedProduct))
            }
          )
      }
    }

    def delete(id: Int): Action[AnyContent] = Action {
      DB.products.indexWhere(_.id == id) match {
        case -1 => NotFound(Json.obj("error" -> s"Product $id not found"))
        case idx =>
          DB.products.remove(idx)
          NoContent // 204
      }
    }
}