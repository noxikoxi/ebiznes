package models

import scala.collection.mutable.ListBuffer

// object -> jest zawsze singletonem
object DB {
    val products : ListBuffer[Product] = ListBuffer(
    Product(1, "Laptop", 2500.0, 1),
    Product(2, "Telefon", 1500.0, 1),
    Product(3, "Miś 150cm", 20.0, 2),
    Product(4, "Czerwone autko pilot", 250.0, 2),
    Product(5, "Woda gazowana 6x1.5L", 15.0, 3),
    Product(6, "Chałwa ", 5.0, 3),
    )

    val categories : ListBuffer[Category] = ListBuffer(
      Category(1, "Elektronika"),
      Category(2, "Zabawki"),
      Category(3, "Żywność")
    )

    val carts: ListBuffer[Cart] = ListBuffer(
      Cart(1, ListBuffer((1, 2), (2, 2))),
      Cart(2, ListBuffer((5, 10), (6, 3))),
    )

    var nextProductId: Int = 7;
    var nextCategoryId: Int = 4;
    var nextCartId: Int = 3;
}
