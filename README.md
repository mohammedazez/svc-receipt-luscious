1. Buatlah struktur database (ERD dan/atau Table Relationship). (20)
![ERD Aplikasi Receipt Luscious drawio](https://github.com/mohammedazez/svc-receipt-luscious/assets/37678093/6cec7cc4-c5ec-4d73-989f-35fc8afeb88f)
![Example data drawio](https://github.com/mohammedazez/svc-receipt-luscious/assets/37678093/134dd1e4-dd48-4d61-8427-fb332b8662f3)
![User Interface](https://github.com/mohammedazez/svc-receipt-luscious/assets/37678093/9e79dca2-68aa-427d-a47e-a99c8a61ce21)








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
       "message": "OK",
       "code": 200,
       "data": [
         {
           "id": "1b2457bf-f52e-4ce0-8c42-a87952ef0085",
           "ingredient_name": "cabe",
           "recipe_id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "quantity": "2 siung",
           "created_at": "2023-07-16 12:28:54.652633035 +0700 WIB m=+33.920682192",
           "updated_at": ""
         },
         {
           "id": "7d19d7fb-a2cd-43fe-997e-a7a7762af1a0",
           "ingredient_name": "telur",
           "recipe_id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "quantity": "2 siung",
           "created_at": "2023-07-16 12:29:01.405170106 +0700 WIB m=+40.673219263",
           "updated_at": ""
         },
         {
           "id": "94278d4b-ee03-4fa3-8611-f3b36ece8c03",
           "ingredient_name": "bawang putih",
           "recipe_id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "quantity": "2 siung",
           "created_at": "2023-07-16 12:28:32.2856108 +0700 WIB m=+11.553660027",
           "updated_at": ""
         },
         {
           "id": "b974800f-570a-4131-b124-a8db262b17cc",
           "ingredient_name": "Bawang Merah",
           "recipe_id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "quantity": "2 siung",
           "created_at": "2023-07-16 12:27:30.293262911 +0700 WIB m=+220.170096059",
           "updated_at": ""
         },
         {
           "id": "c402b3ba-fb99-4d4c-9ca6-b6c4bbd7fdef",
           "ingredient_name": "garam",
           "recipe_id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "quantity": "2 siung",
           "created_at": "2023-07-16 12:28:49.597249515 +0700 WIB m=+28.865298672",
           "updated_at": ""
         },
         {
           "id": "c80cc89d-511c-43ef-8aa3-446f829e9766",
           "ingredient_name": "Bawang Merah",
           "recipe_id": "643f8d13a789904bbfd7c3b3",
           "quantity": "2 siung",
           "created_at": "2023-07-16 11:45:40.23911817 +0700 WIB m=+3389.297679727",
           "updated_at": ""
         },
         {
           "id": "ee58f1b9-ed8c-4acd-a78a-57e32f2760e8",
           "ingredient_name": "ketumbar",
           "recipe_id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "quantity": "2 siung",
           "created_at": "2023-07-16 12:28:58.314253859 +0700 WIB m=+37.582303086",
           "updated_at": ""
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
       "message": "Created",
       "code": 201
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
       "message": "Bad Request",
       "code": 400,
       "error_validation": {
         "ingredient_name": "Nama bahan makanan harus diisi",
         "quantity": "kuantitas harus diisi",
         "recipe_id": "recipe id harus diisi"
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
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "ingredient_name": "Bawang Merah",
         "recipe_id": "643f8b755823cf72b856117b",
         "quantity": "4 siung",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": "2023-07-16T06:41:23.451Z"
       }
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
       "message": "OK",
       "code": 200,
       "data": [
         {
           "id": "594dbadc-515a-4c10-871e-ec8c31dd0d83",
           "category_name": "Menu Utama",
           "created_at": "2023-07-16 11:46:42.61695346 +0700 WIB m=+3451.675515087",
           "updated_at": ""
         },
         {
           "id": "6e9cc8a1-8e4f-4a49-a902-5bdcef113f22",
           "category_name": "Masakan Sayuran",
           "created_at": "2023-07-16 12:04:58.244700473 +0700 WIB m=+118.675945563",
           "updated_at": ""
         },
         {
           "id": "6eaf678e-605d-4a84-9011-54ab294add5a",
           "category_name": "Masakan Ikan",
           "created_at": "2023-07-16 12:05:06.032098447 +0700 WIB m=+126.463343048",
           "updated_at": ""
         },
         {
           "id": "79424f11-f199-4f3d-9f33-34718ac125c3",
           "category_name": "Hidangan Penutup",
           "created_at": "2023-07-16 12:05:19.215896749 +0700 WIB m=+139.647141350",
           "updated_at": ""
         },
         {
           "id": "7ac65388-7762-489a-90bd-508ec6a117b8",
           "category_name": "Hidangan Pembuka",
           "created_at": "2023-07-16 11:54:55.800392091 +0700 WIB m=+2.845178231",
           "updated_at": ""
         },
         {
           "id": "99f47b6a-a321-4bca-8767-a1cebf12c486",
           "category_name": "Minuman",
           "created_at": "2023-07-16 12:05:12.806207917 +0700 WIB m=+133.237452517",
           "updated_at": ""
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
     
   - `GET /api/v1/category/:category_id 	`
     Method ini akan digunakan untuk menampilkan detail kategori makanan.
   
     > Get detail category
   
     
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category/643f8d13a789904bbfd7c3b3
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "category_name": "Masakan Ikan",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": null,
         "recipe": [
           {
             "id": "643f8d13a789904bbfd7c3b3",
             "recipe_name": "Ikan Mas Bakar",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8ccba789904bbfd7c3b0",
             "recipe_name": "Ikan Gurame Goreng",
             "created_at": "2023-07-15T06:41:23.451Z"
           },
           {
             "id": "643f8c6b0f2d077622e36292",
             "recipe_name": "Pepes Ikan",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8c4d0f2d077622e36291",
             "recipe_name": "Ikan Tuna Bakar",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8c3a0f2d077622e36290",
             "recipe_name": "Ikan Asam Manis",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8bd85823cf72b856117c5",
             "recipe_name": "Ikan Tuna Bakar",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           }
         ]
       }
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
       "message": "OK",
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
     
     Bad Request (400 - bad request)
     
     ```json
     {
       "status": false,
       "message": "Bad Request",
       "code": 400,
       "error_validation": {
         "category_name": "kategory makanan harus diisi"
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
       "message": "OK",
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


   - `DELETE /api/v1/category/:category_id`

     Method ini akan digunakan untuk menghapus data category.

     > Delete category
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category/643f8d13a789904bbfd7c3b3
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
     Method ini akan digunakan untuk menampilkan resep makanan.

     > Get All recipe

     
   
     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "OK",
       "code": 200,
       "data": [
         {
           "id": "32df356c-cd05-4f5c-89ad-6b9c4101e028",
           "recipe_name": "ikan gurame goreng",
           "created_at": "2023-07-16 12:17:58.068777226 +0700 WIB m=+137.648063193",
           "updated_at": ""
         },
         {
           "id": "35ca399a-1ba2-4a1a-a5e0-b8db89797f5a",
           "recipe_name": "ikan tuna bakar",
           "created_at": "2023-07-16 12:19:24.029849527 +0700 WIB m=+1.843695320",
           "updated_at": ""
         },
         {
           "id": "3f3f1126-adf3-468b-b732-1d457ba469dd",
           "recipe_name": "ikan tuna bakar",
           "created_at": "2023-07-16 12:18:11.972269787 +0700 WIB m=+151.551555265",
           "updated_at": ""
         },
         {
           "id": "3fed7fad-bc8a-4e17-b9bd-8950bd278c38",
           "recipe_name": "ikan tuna bakar",
           "created_at": "2023-07-16 12:19:07.064172264 +0700 WIB m=+206.643458301",
           "updated_at": ""
         },
         {
           "id": "96c5778f-a995-4bcb-9445-702a595b6f12",
           "recipe_name": "ikan asam manis",
           "created_at": "2023-07-16 12:18:17.97036506 +0700 WIB m=+157.549650538",
           "updated_at": ""
         },
         {
           "id": "d122bd17-01e3-4764-8553-1ea34079a5e2",
           "recipe_name": "ikan tuna bakar",
           "created_at": "2023-07-16 12:19:28.755417643 +0700 WIB m=+6.569263925",
           "updated_at": ""
         },
         {
           "id": "e898f8df-a661-405b-b9f7-20b72a7a4704",
           "recipe_name": "ikan mas bakar",
           "created_at": "2023-07-16 12:17:37.147500103 +0700 WIB m=+116.726785651",
           "updated_at": ""
         },
         {
           "id": "f511b1da-b1e9-4df4-bc52-660f73f407f8",
           "recipe_name": "pepes ikan",
           "created_at": "2023-07-16 12:18:05.068311764 +0700 WIB m=+144.647597312",
           "updated_at": ""
         },
         {
           "id": "f812a69b-364c-4513-bfed-fef24c951dd0",
           "recipe_name": "ikan tuna bakar",
           "created_at": "2023-07-16 12:18:24.438308749 +0700 WIB m=+164.017594297",
           "updated_at": ""
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


- `GET /api/v1/recipe/:recipe_id`
     Method ini akan digunakan untuk menampilkan detail resep.
   
     > Get detail recipe
   
     
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe/643f8d13a789904bbfd7c3b3
     ```
   
     _Response (200)_
   
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "id": "643f8d13a789904bbfd7c3b3",
         "recipe_name": "Ikan Tuna Bakar",
         "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
         "created_at": "2023-07-15T06:41:23.451Z",
         "updated_at": null,
         "ingredient": [
           {
             "id": "643f8d13a789904bbfd7c3b3",
             "ingredient_name": "Bawang Merah",
             "recipe_id": "643f8d13a789904bbfd7c3b3",
             "quantity": "2 siung",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8ccba789904bbfd7c3b0",
             "ingredient_name": "Bawang Putih",
             "recipe_id": "643f8d13a789904bbfd7c3b3",
             "quantity": "2 siung",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           },
           {
             "id": "643f8c6b0f2d077622e36292",
             "ingredient_name": "Garam",
             "recipe_id": "643f8d13a789904bbfd7c3b3",
             "quantity": "secukupnya",
             "created_at": "2023-07-15T06:41:23.451Z",
             "updated_at": null
           }
         ]
       }
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
     
- `POST /api/v1/recipe`
       Method ini akan digunakan untuk menambahkan resep makanan.

     > Add new recipe

     

     _Request Body_

     ```json
     {
       "recipe_name": "Menu Utama",
       "category_id": "643f8c4d0f2d077622e36291",
       "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
       "ingredient": [
         {
           "id": "643f8d13a789904bbfd7c3b3"
         },
         {
           "id": "643f8ccba789904bbfd7c3b0"
         },
         {
           "id": "643f8c6b0f2d077622e36292"
         }
       ]
     }
     ```
   
     _Response (201)_
   
     ```json
     {
       "status": true,
       "message": "created",
       "code": 201,
       "data": {
         "recipe_name": "Menu Utama",
         "category_id": "643f8c4d0f2d077622e36291",
         "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
         "ingredient": [
           {
             "id": "643f8d13a789904bbfd7c3b3"
           },
           {
             "id": "643f8ccba789904bbfd7c3b0"
           },
           {
             "id": "643f8c6b0f2d077622e36292"
           }
         ]
       }
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
       "message": "Bad Request",
       "code": 400,
       "error_validation": {
         "category_id": "kategori makanan harus diisi",
         "how_to_make": "cara pembuatan harus diisi",
         "recipe_name": "judul resep makanan harus diisi"
       }
     }
     ```
   
   ---
   
- `PUT /api/v1/recipe/:recipe_id`
     Method ini akan digunakan untuk update data resep makanan.

     > Edit resep makanan

     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe/643f8d13a789904bbfd7c3b3
     ```
     
     _Request Body_
     
     ```json
     {
       "recipe_name": "Menu Utama",
       "category_id": "643f8c4d0f2d077622e36291",
       "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
       "ingredient": [
         {
           "id": "643f8d13a789904bbfd7c3b3"
         },
         {
           "id": "643f8ccba789904bbfd7c3b0"
         },
         {
           "id": "643f8c6b0f2d077622e36292"
         }
       ]
     }
     ```
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": {
         "recipe_name": "Menu Utama",
         "category_id": "643f8c4d0f2d077622e36291",
         "how_to_make": "Kupas kentang, lalu potong dadu dan goreng hingga matang. Sisihkan. kemudian bilas ikan teri lalu goreng hingga matang. Sisihkan. Panaskan sedikit minyak lalu tumis bawang merah dan bawang putih hingga harum. Kemudian masukkan cabai dan tumis hingga harum. Masukkan kentang, ikan teri, dan daun bawang. Lalu bumbui dengan garam, gula dan kaldu bubuk secukupnya. Aduk rata. Jangan lupa koreksi rasa, jika sudah pas angkat dan siap disajikan.",
         "ingredient": [
           {
             "id": "643f8d13a789904bbfd7c3b3"
           },
           {
             "id": "643f8ccba789904bbfd7c3b0"
           },
           {
             "id": "643f8c6b0f2d077622e36292"
           }
         ]
       }
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
     
- `DELETE /api/v1/recipe/:recipe_id`
     Method ini akan digunakan untuk menghapus data resep makanan.

     > Delete recipe
   
     _Path Example_
   
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/recipe/643f8d13a789904bbfd7c3b3
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


4. Membuat API CRUD Menampilkan list / index / read resep berdasarkan search / filter bahan dan kategori yang digunakan. (20)

     Filter Bahan makanan

     _Path Example_
     
     ```
     https://services.app.com/svc-receipt-luscious/api/v1/ingredient?ingredient_name=garam
     ```

     | Parameter       | Data Type | Location |
     | :-------------- | :-------: | -------: |
     | ingredient_name |  String   |    query |
     
     _Response (200)_
     
     ```json
     {
       "status": true,
       "message": "ok",
       "code": 200,
       "data": [
         {
           "id": "643f8c6b0f2d077622e36292",
           "ingredient_name": "Garam",
           "recipe_id": "643f8b755823cf72b856117b",
           "quantity": "secukupnya,
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
     
     

     
     
     Filter kategori makanan
     
     _Path Example_

     ```
     https://services.app.com/svc-receipt-luscious/api/v1/category?category_name=menu
     ```
     
     | Parameter     | Data Type | Location |
     | :------------ | :-------: | -------: |
     | category_name |  String   |    query |
     
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
           "created_at": "2023-07-15T06:41:23.451Z",
           "updated_at": null
         }
       ]
     }
     ```
     


5. Buat unit test untuk setiap fungsi yang telah dikerjakan (20)
