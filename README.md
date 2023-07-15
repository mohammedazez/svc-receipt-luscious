1. Buatlah struktur database (ERD dan/atau Table Relationship). (20)
![ERD Aplikasi Receipt Luscious drawio](https://github.com/mohammedazez/svc-receipt-luscious/assets/37678093/43915351-a1ec-4d6d-b472-98043b1458f6)
![Example data](https://github.com/mohammedazez/svc-receipt-luscious/assets/37678093/2fc2376a-5a19-4a74-b4a9-6afcf7246416)
![User Interface](https://github.com/mohammedazez/svc-receipt-luscious/assets/37678093/fedf7b13-d10b-432c-8f47-eb0fd3d27cfd)




3. Membuat API CRUD (create, read / index, update, delete) data master (bahan dan kategori). (20)

   

   ### Ingredients

   - Endpoint: `GET /api/v1/ingredient	`

     Method ini akan digunakan untuk menampilkan bahan makanan.

     > Get All Ingredients

     

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient
     ```

     _Response (200)_

     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "ingredient_name": "Bawang Merah",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "2 siung",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8ccba789904bbfd7c3b0",
           "ingredient_name": "Bawang Putih",
           "recipe_id": "643f8a7ebc2985b6227584ee",
           "quantity": "2 siung",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8c6b0f2d077622e36292",
           "ingredient_name": "Garam",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "secukupnya,
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8c4d0f2d077622e36291",
           "ingredient_name": "Cabe",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "secukupnya",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8c3a0f2d077622e36290",
           "ingredient_name": "Ketumbar",
           "recipe_id": "643f8a7ebc2985b6227584ee",
           "quantity": "secukupnya",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8bd85823cf72b856117c5",
           "ingredient_name": "Telur",
           "recipe_id": "643f8d13a789904bbfd7c3b3",
           "quantity": "2 buah",
           "created_at": "2023-07-15T06:41:23.451Z"
         }
       ]
     }
     ```

     _Response (500 - Internal Server Error)_

     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```

     ---

   - `POST /api/v1/ingredient`
     Method ini akan digunakan untuk menambahkan data bahan makanan.

     > Add new ingredient

     

     _Request Body_

     ```json
     {
       "ingredient_name": "Bawang Merah",
       "recipe_id": "643f8d13a789904bbfd7c3b3",
       "quantity": "2 siung"
     }
     ```
     
     _Response (201)_
     
     ```json
     {
       "status": true,
       "message": "created",
       "code": 201,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "ingredient_name": "Bawang Merah",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "2 siung",
           "created_at": "2023-07-15T06:41:23.451Z"
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     Bad Request (400 - bad request)
     
     ```json
     {
       "status": false,
       "message": "bad request",
       "code": 400,
       "error": {
         "ingredient_name": "cannot be blank",
         "recipe_id": "cannot be blank",
         "quantity": "cannot be blank"
       }
     }
     ```
     
     ---
     
     
     
   - `PUT /api/v1/ingredient/:ingredient_id`
     Method ini akan digunakan untuk update data bahan makanan.

     > Edit ingredient

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient/643f8d13a789904bbfd7c3b3
     ```
     
     _Request Body_
     
     ```json
     {
       "ingredient_name": "Bawang Merah",
       "recipe_id": "643f8d13a789904bbfd7c3b3",
       "quantity": "4 siung"
     }
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "ingredient_name": "Bawang Merah",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "4 siung",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": "2023-07-16T06:41:23.451Z"
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
   - `DELETE /api/v1/ingredients/:ingredient_id`
     Method ini akan digunakan untuk menghapus data bahan makanan.

     > Delete ingredient
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient/643f8d13a789904bbfd7c3b3
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---
     
     
   
   ### Category
   
   - `GET /api/v1/category	`
     Method ini akan digunakan untuk menampilkan kategori makanan.

     > Get All category

     
   
     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "category_name": "Menu Utama",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8ccba789904bbfd7c3b0",
           "category_name": "Hidangan Pembuka",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8c6b0f2d077622e36292",
           "category_name": "Masakan Sayuran",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8c4d0f2d077622e36291",
           "category_name": "Masakan Ikan",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8c3a0f2d077622e36290",
           "category_name": "Minuman",
           "created_at": "2023-07-15T06:41:23.451Z"
         },
         {
           "id": "643f8bd85823cf72b856117c5",
           "category_name": "Hidangan Penutup",
           "created_at": "2023-07-15T06:41:23.451Z"
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---


   - `POST /api/v1/category`
       Method ini akan digunakan untuk menambahkan data kategori makanan.

     > Add new category

     

     _Request Body_

     ```json
     {
       "category_name": "Menu Utama"
     }
     ```
     
     _Response (201)_
     
     ```json
     {
       "status": true,
       "message": "created",
       "code": 201,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "category_name": "Menu Utama",
           "created_at": "2023-07-15T06:41:23.451Z"
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     Bad Request (400 - bad request)
     
     ```json
     {
       "status": false,
       "message": "bad request",
       "code": 400,
       "error": {
         "Menu Utama": "cannot be blank"
       }
     }
     ```
     
     ---



   - `PUT /api/v1/category/:category_id`
     Method ini akan digunakan untuk update data category.

     > Edit category

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category/643f8d13a789904bbfd7c3b3
     ```
     
     _Request Body_
     
     ```json
     {
       "category_name": "Menu Utama Baru"
     }
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8d13a789904bbfd7c3b3",
           "category_name": "Menu Utama Baru",
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": "2023-07-16T06:41:23.451Z"
         }
       ]
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---


   - `DELETE /api/v1/category/:category_id`

     Method ini akan digunakan untuk menghapus data category.

     > Delete category
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/cateogry/643f8d13a789904bbfd7c3b3
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200
     }
     ```
     
     _Response (500 - Internal Server Error)_
     
     ```json
     {
       "status": false,
       "message": "internal server error",
       "code": 500
     }
     ```
     
     ---




2. Membuat API CRUD data resep dan beserta menampilkan bahan-bahan dan kategori yang digunakan. (20)

### 	Recipe

- `GET /api/v1/recipe	`
- `POST /api/v1/recipe`
- `PUT /api/v1/recipe/:recipe_id`
- `DELETE /api/v1/recipe/:recipe_id`

4. Membuat API CRUD Menampilkan list / index / read resep berdasarkan search / filter bahan dan kategori yang digunakan. (20)

5. Buat unit test untuk setiap fungsi yang telah dikerjakan (20)
