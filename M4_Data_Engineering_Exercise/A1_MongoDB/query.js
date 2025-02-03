//1. Find High-Spending Users

db.orders.aggregate([
    {
      $group: {
        _id: "$userId",
        totalSpent: { $sum: "$totalAmount" }
      }
    },
    {
      $match: { totalSpent: { $gt: 500 } }
    },
    {
      $lookup: {
        from: "users",
        localField: "_id",
        foreignField: "userId",
        as: "userDetails"
      }
    },
    {
      $unwind: "$userDetails"
    },
    {
      $project: {
        userId: "$_id",
        name: "$userDetails.name",
        email: "$userDetails.email",
        totalSpent: 1
      }
    }
  ]);

  
  //2. List Popular Products by Average Rating

  db.products.aggregate([
    { $unwind: "$ratings" },
    {
      $group: {
        _id: "$productId",
        avgRating: { $avg: "$ratings.rating" },
        name: { $first: "$name" }
      }
    },
    {
      $match: { avgRating: { $gte: 4 } }
    },
    {
      $project: {
        productId: "$_id",
        name: 1,
        avgRating: 1
      }
    }
  ]);

  
  //3. Search for Orders in a Specific Time Range

  db.orders.aggregate([
    {
      $match: {
        orderDate: {
          $gte: ISODate("2024-12-01T00:00:00Z"),
          $lte: ISODate("2024-12-31T23:59:59Z")
        }
      }
    },
    {
      $lookup: {
        from: "users",
        localField: "userId",
        foreignField: "userId",
        as: "userDetails"
      }
    },
    {
      $unwind: "$userDetails"
    },
    {
      $project: {
        orderId: 1,
        orderDate: 1,
        totalAmount: 1,
        status: 1,
        userName: "$userDetails.name"
      }
    }
  ]);

  
  //4. Update Stock After Order Completion

  db.orders.find({ status: "Delivered" }).forEach(order => {
    order.items.forEach(item => {
      db.products.updateOne(
        { productId: item.productId },
        { $inc: { stock: -item.quantity } }
      );
    });
  });

  
  //5. Find Nearest Warehouse

  db.warehouses.aggregate([
    {
      $geoNear: {
        near: { type: "Point", coordinates: [-74.006, 40.7128] },
        distanceField: "distance",
        maxDistance: 50000, // 50 kilometers in meters
        query: { products: "P001" },
        spherical: true
      }
    },
    {
      $project: {
        warehouseId: 1,
        distance: 1,
        location: 1
      }
    }
  ]);
  