package controllers

import javax.inject.*
import play.api.libs.json.*
import play.api.mvc.{Action, *}
import models.{Cart, DB, Product, CartData, CartItemData}


@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  def getAll: Action[AnyContent] = Action {
    val cartsWithProductsNames = DB.carts.map { cart =>
      Json.obj(
        "cart_id" -> cart.id,
        "items" -> cart.products.map { case (product_id, quantity) =>
          val productName = DB.products.find(_.id == product_id).map(_.name).getOrElse("Unknown")
          Json.obj(
            "name" -> productName,
            "quantity" -> quantity
          )
        }
      )
    }

    Ok(Json.toJson(cartsWithProductsNames))
  }

  def getById(id: Int): Action[AnyContent] = Action {
    DB.carts.find(_.id == id) match {
      case Some(cart) => Ok(Json.obj(
      "cart_id" -> cart.id,
            "items" -> cart.products.map { case (product_id, quantity) =>
              val productName = DB.products.find(_.id == product_id).map(_.name).getOrElse("Unknown")
              Json.obj(
                "name" -> productName,
                "quantity" -> quantity
              )
            }
          )
      )
      case None => NotFound(Json.obj("error" -> s"Category $id not found"))
    }
  }

  def add: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartData].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid cart format")),
      cartData => {
        val newCart = Cart(DB.nextCartId, cartData.products)
        DB.nextCartId += 1
        DB.carts += newCart
        Created(Json.toJson(newCart))
      }
    )
  }

  def updateItem(id: Int): Action[JsValue] = Action(parse.json) { request =>
    DB.carts.find(_.id == id) match {
      case Some(cart) => request.body.validate[CartItemData].fold(
          errors => BadRequest(Json.obj("error" -> "Invalid cart item format")),
          cartItemData => {
            cart.products.indexWhere(_._1 == cartItemData.product_id) match {
              case -1 =>
                // Produkt nie istnieje
                cart.products += ((cartItemData.product_id, cartItemData.quantity))

              case idx =>
                // Produkt istnieje, aktualizujemy ilość
                val (product_id, oldQuantity) = cart.products(idx)
                cart.products.update(idx, (product_id, oldQuantity + cartItemData.quantity))
            }
            Created(Json.toJson(cart))
          }
        )


      case None => NotFound(Json.obj("error" -> s"Cart $id not found"))
    }
  }


  def delete(id: Int): Action[AnyContent] = Action {
    DB.carts.indexWhere(_.id == id) match {
      case -1 => NotFound(Json.obj("error" -> s"Cart $id not found"))
      case idx =>
        DB.carts.remove(idx)
        NoContent
    }
  }
}
