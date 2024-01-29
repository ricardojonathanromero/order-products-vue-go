catalogsDB = db.getSiblingDB('catalogs');

catalogsDB.createCollection('products', {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            title: "Admins Object Validation",
            required: [ "sku", "price", "name", "quantity" ],
            properties: {
                _id: {
                    bsonType: "objectId",
                    description: "primary key"
                },
                sku: {
                    bsonType: "string",
                    description: "'sku' must be a string and is required"
                },
                price: {
                    bsonType: "double",
                    description: "'price' must be a float64 and is required"
                },
                name: {
                    bsonType: "string",
                    description: "'name' must be a string and is required"
                },
                description: {
                    bsonType: "string",
                    description: "'description' must be a string"
                },
                quantity: {
                    bsonType: "int64",
                    description: "'quantity' must be an int64 and is required"
                },
                image: {
                    bsonType: "string",
                    description: "'image' must be a string"
                },
                categories: {
                    bsonType: "array",
                    description: "'categories' must be an array of strings",
                    items: {
                        bsonType: "string",
                        description: "'item' must be a string"
                    }
                }
            }
        }
    }
});

catalogsDB.products.insertMany(
    [
        {
            sku: "75007614",
            price: 18.30,
            name: "Coca Cola 600ml",
            description: "Coca Cola 600ml",
            quantity: 1000,
            image: "https://elmundoderikachito.com.mx/wp-content/uploads/2022/03/coca-600ml-100x100.jpg",
            categories: ["soda"]
        },
        {
            sku: "75007615",
            price: 15.60,
            name: "Refresco Coca Cola 500 ml",
            description: "Refresco Coca Cola 500 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00000007500980L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251906",
            price: 31.00,
            name: "Jugo de Naranja",
            description: "Jugo Del Valle de naranja 100% jugo 946 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105535957L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251907",
            price: 31.00,
            name: "Jugo de Manzana",
            description: "Jugo Del Valle 100% de manzana 946 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105535776L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251908",
            price: 29,
            name: "Fanta 2L",
            description: "Refresco Fanta sabor naranja de 2 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105530387L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251909",
            price: 29,
            name: "Fresca 2L",
            description: "Refresco Fresca sabor toronja de 2 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105530388L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251910",
            price: 17.00,
            name: "Sidral Mundet 600ml",
            description: "Refresco Sidral Mundet sabor manzana de 600 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105533998L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251911",
            price: 29.00,
            name: "Refresco Coca Cola sin azúcar 2.5 l",
            description: "Refresco Coca Cola sin azúcar 2.5 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105532168L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251912",
            price: 118.00,
            name: "Refresco Coca Cola sin azúcar mini latas 8 pack de 235 ml c/u",
            description: "Refresco Coca Cola sin azúcar mini latas 8 pack de 235 ml c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105534991L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251913",
            price: 18.00,
            name: "Agua mineral Topo Chico 600 ml",
            description: "Agua mineral Topo Chico 600 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00002113601054L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["soda"]
        },
        {
            sku: "10251914",
            price: 45.00,
            name: "Aceite puro de soya Nutrioli 946 ml",
            description: "Aceite puro de soya Nutrioli 946 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750103912014L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251915",
            price: 46.00,
            name: "Aceite vegetal 123 1 l",
            description: "Aceite vegetal 123 1 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00000007500234L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251916",
            price: 19.00,
            name: "Arroz Great Value super extra 900 g",
            description: "Arroz Great Value super extra 900 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179163863L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251917",
            price: 129.00,
            name: "Café soluble Great Value colombiano gourmet 200 g",
            description: "Café soluble Great Value colombiano gourmet 200 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101040785L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251918",
            price: 167.00,
            name: "Sustituto de crema para café Coffee Mate en polvo 1.2 kg",
            description: "Sustituto de crema para café Coffee Mate en polvo 1.2 kg",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750647510354L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251919",
            price: 47.50,
            name: "Pan blanco Bimbo grande 680 g",
            description: "Pan blanco Bimbo grande 680 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750100011120L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251920",
            price: 39.00,
            name: "Tortillas de harina Tía Rosa Tortillinas 561 g",
            description: "Tortillas de harina Tía Rosa Tortillinas 561 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750103048666L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251921",
            price: 32.00,
            name: "Pan tostado Bimbo clásico 210 g",
            description: "Pan tostado Bimbo clásico 210 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750100011180L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251922",
            price: 13.00,
            name: "Atún Aurrera en agua 130 g",
            description: "Atún Aurrera en agua 130 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750649500405L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251923",
            price: 69.00,
            name: "Azúcar Great Value refinada 2 kg",
            description: "Azúcar Great Value refinada 2 kg",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179162335L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["groceries"]
        },
        {
            sku: "10251924",
            price: 24.50,
            name: "Pasta Barilla spaghetti No.5 500 g",
            description: "Pasta Barilla spaghetti No.5 500 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00807680951520L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251925",
            price: 12.00,
            name: "Spaghetti La Moderna 250 g",
            description: "Spaghetti La Moderna 250 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101832203L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251926",
            price: 7.50,
            name: "Puré de tomate La Costeña condimentado 210 g",
            description: "Puré de tomate La Costeña condimentado 210 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101700561L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251927",
            price: 91.00,
            name: "Pasta Barilla lasagne 500 g",
            description: "Pasta Barilla lasagne 500 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00807680957567L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251928",
            price: 43.00,
            name: "Crema Campbells de elote 430 g",
            description: "Crema Campbells de elote 430 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101136149L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251929",
            price: 12.00,
            name: "Sopa de codo La Moderna N.2, 250 g",
            description: "Sopa de codo La Moderna N.2, 250 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101832200L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251930",
            price: 43.00,
            name: "Crema Campbells de espárragos 420 g",
            description: "Crema Campbells de espárragos 420 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101131202L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251931",
            price: 43.00,
            name: "Crema Campbells de champiñones 420 g",
            description: "Crema Campbells de champiñones 420 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101131262L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251932",
            price: 43.00,
            name: "Spaghetti La Moderna 1 kg",
            description: "Spaghetti La Moderna 1 kg",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101831140L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251933",
            price: 43.00,
            name: "Sopa Campbells lentejas a la mexicana 430 g",
            description: "Sopa Campbells lentejas a la mexicana 430 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750101136145L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["pasta"]
        },
        {
            sku: "10251934",
            price: 14.00,
            name: "Sopa instantánea Maruchan Instant Lunch con camarón y chile piquín 64 g",
            description: "Sopa instantánea Maruchan Instant Lunch con camarón y chile piquín 64 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00004178900198L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251935",
            price: 14.00,
            name: "Sopa instantánea Maruchan Instant Lunch con camarón limón y habanero 64 g",
            description: "Sopa instantánea Maruchan Instant Lunch con camarón limón y habanero 64 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00004178900186L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251936",
            price: 36.00,
            name: "Arroz Knorr a la mexicana 155 g",
            description: "Arroz Knorr a la mexicana 155 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750630630720L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251937",
            price: 35.00,
            name: "Arroz Verde Valle precocido blanco 300 g",
            description: "Arroz Verde Valle precocido blanco 300 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750107130266L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251938",
            price: 36.00,
            name: "Arroz Knorr primavera 155 g",
            description: "Arroz Knorr primavera 155 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750100511428L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251939",
            price: 62.00,
            name: "Chilorio Chata de cerdo 215 g",
            description: "Chilorio Chata de cerdo 215 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750102350712L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251940",
            price: 64.00,
            name: "Chilorio Chata de pavo 215 g",
            description: "Chilorio Chata de pavo 215 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750102350713L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251941",
            price: 62.00,
            name: "Caldo de pollo Gallina Blanca casero 1 l",
            description: "Caldo de pollo Gallina Blanca casero 1 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00841030034905L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251942",
            price: 39.00,
            name: "Chilaquiles La Sierra con salsa verde 370 g",
            description: "Chilaquiles La Sierra con salsa verde 370 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105246201L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251943",
            price: 18.00,
            name: "Frijoles negros Isadora refritos 430 g",
            description: "Frijoles negros Isadora refritos 430 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750107130779L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["instant_foods"]
        },
        {
            sku: "10251944",
            price: 140.00,
            name: "Leche Lala entera 6 pzas 1 l c/u",
            description: "Leche Lala entera 6 pzas 1 l c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750102056596L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251945",
            price: 89.00,
            name: "Yoghurt bebible Alpura multi sabor 10 pzas 220 g c/u",
            description: "Yoghurt bebible Alpura multi sabor 10 pzas 220 g c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105591495L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251946",
            price: 85.00,
            name: "Yoghurt griego Oikos natural monkfruit con 11 g de proteína 900 g",
            description: "Yoghurt griego Oikos natural monkfruit con 11 g de proteína 900 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750644310591L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251947",
            price: 95.00,
            name: "Alimento lácteo fermentado Activia ciruela pasa y fresa sin azúcar 8 pzas de 225 g c/u",
            description: "Alimento lácteo fermentado Activia ciruela pasa y fresa sin azúcar 8 pzas de 225 g c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750644310356L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251948",
            price: 10.00,
            name: "Leche Santa Clara sabor helado de fresa 180 ml",
            description: "Leche Santa Clara sabor helado de fresa 180 ml",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105537719L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251949",
            price: 69.00,
            name: "Yoghurt Danone con fresa 8 pzas de 220 g c/u",
            description: "Yoghurt Danone con fresa 8 pzas de 220 g c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750103239868L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251950",
            price: 150,
            name: "Leche Santa Clara entera caja con 6 pzas de 1 l c/u",
            description: "Leche Santa Clara entera caja con 6 pzas de 1 l c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105536750L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251951",
            price: 37.00,
            name: "Leche Lala 100 sin lactosa proteína light 1 l",
            description: "Leche Lala 100 sin lactosa proteína light 1 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750102055097L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251952",
            price: 87.00,
            name: "Huevo blanco San Juan 30 pzas",
            description: "Huevo blanco San Juan 30 pzas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750300055517L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251953",
            price: 56.50,
            name: "Huevo blanco San Juan 18 pzas",
            description: "Huevo blanco San Juan 18 pzas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750300055509L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["dairy"]
        },
        {
            sku: "10251954",
            price: 159.50,
            name: "Salmón Marketside sin piel 500 g",
            description: "Salmón Marketside sin piel 500 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750649501091L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251955",
            price: 189.00,
            name: "Arrachera de res marinada SuKarne 600 g",
            description: "Arrachera de res marinada SuKarne 600 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750129350025L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251956",
            price: 184.00,
            name: "Arrachera de res Marketside marinada 600 g",
            description: "Arrachera de res Marketside marinada 600 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179166302L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251957",
            price: 195.00,
            name: "Milanesa de Pechuga de Pollo por kilo",
            description: "Milanesa de Pechuga de Pollo por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020196100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251958",
            price: 99.00,
            name: "Filete de pescado Sierra Madre Blanco del Nilo pimienta limón 500 g",
            description: "Filete de pescado Sierra Madre Blanco del Nilo pimienta limón 500 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00069402113700L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251959",
            price: 84.00,
            name: "Filete de tilapia por kilo",
            description: "Filete de tilapia por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020608100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251960",
            price: 319.00,
            name: "Camarón grande sin cabeza por kg",
            description: "Camarón grande sin cabeza por kg",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020612100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251961",
            price: 189.00,
            name: "Camarón coctelero por kilo",
            description: "Camarón coctelero por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020644500000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251962",
            price: 74.00,
            name: "Filete basa oriental rojo por kilo",
            description: "Filete basa oriental rojo por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020610600000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251963",
            price: 249.00,
            name: "Medallones de atún Dolores Premium aleta amarilla 4 piezas",
            description: "Medallones de atún Dolores Premium aleta amarilla 4 piezas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750104540212L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["meat", "fishes"]
        },
        {
            sku: "10251964",
            price: 185.00,
            name: "Papel higiénico Great Value premium 32 rollos con 220 hojas dobles c/u",
            description: "Papel higiénico Great Value premium 32 rollos con 220 hojas dobles c/u",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179165547L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251965",
            price: 169.00,
            name: "Detergente líquido Great Value para ropa de color 7 l",
            description: "Detergente líquido Great Value para ropa de color 7 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179165713L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251966",
            price: 52.00,
            name: "Servilletas Great Value premium 400 pzas",
            description: "Servilletas Great Value premium 400 pzas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179166600L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251967",
            price: 35.00,
            name: "Bolsas para basura Great Value Terra mediana 20 pzas",
            description: "Bolsas para basura Great Value Terra mediana 20 pzas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179162129L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251968",
            price: 45.00,
            name: "Lavatrastes líquido Great Value aroma limón 1.5 l",
            description: "Lavatrastes líquido Great Value aroma limón 1.5 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179161481L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251969",
            price: 51.00,
            name: "Papel aluminio Great Value 1 rollo con 10 metros",
            description: "Papel aluminio Great Value 1 rollo con 10 metros",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750179162885L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251970",
            price: 37.50,
            name: "Limpiador líquido Fabuloso frescura activa fresca lavanda 2 l",
            description: "Limpiador líquido Fabuloso frescura activa fresca lavanda 2 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750103590534L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251971",
            price: 110.00,
            name: "Limpiador desinfectante Pinol Aromas floral 5.1 l",
            description: "Limpiador desinfectante Pinol Aromas floral 5.1 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750102540136L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251972",
            price: 112.00,
            name: "Desinfectante antibacterial Lysol Crisp Linen en aerosol 354 g",
            description: "Desinfectante antibacterial Lysol Crisp Linen en aerosol 354 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105879688L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251973",
            price: 54.50,
            name: "Blanqueador desinfectante Cloralex El Rendidor 3.75 l",
            description: "Blanqueador desinfectante Cloralex El Rendidor 3.75 l",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00000007500063L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["house"]
        },
        {
            sku: "10251974",
            price: 50.00,
            name: "Pechuga de Pavo Sabori extrafina 250 g",
            description: "Pechuga de Pavo Sabori extrafina 250 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750151849253L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251975",
            price: 61.00,
            name: "Jamón de pierna Sabori extrafino 250 g",
            description: "Jamón de pierna Sabori extrafino 250 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750151849254L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251976",
            price: 37.00,
            name: "Salchicha de Pavo Sabori 500 g",
            description: "Salchicha de Pavo Sabori 500 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750151849232L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251977",
            price: 62.50,
            name: "Salchicha cocktail de pavo Zwan premium 500 g",
            description: "Salchicha cocktail de pavo Zwan premium 500 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750105772019L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251978",
            price: 99.00,
            name: "Torta de Papa por kilo",
            description: "Torta de Papa por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020128700000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251979",
            price: 130.00,
            name: "Pollo rostizado por pza",
            description: "Pollo rostizado por pza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020903200000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251980",
            price: 320.00,
            name: "Pechuga de pavo San Rafael natural por kilo",
            description: "Pechuga de pavo San Rafael natural por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020380500000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251981",
            price: 29.50,
            name: "Salchicha tipo viena FUD 266 g",
            description: "Salchicha tipo viena FUD 266 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750104000975L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251982",
            price: 55.00,
            name: "Jamón de pavo Lala Plenia virginia 450 g",
            description: "Jamón de pavo Lala Plenia virginia 450 g",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750102056462L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251983",
            price: 296.00,
            name: "Jamón de pierna San Rafael real por kilo",
            description: "Jamón de pierna San Rafael real por kilo",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020548100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["cold_meats"]
        },
        {
            sku: "10251984",
            price: 125.00,
            name: "Pizza sabor hawaiana",
            description: "Pizza sabor hawaiana",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020826800000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251985",
            price: 9.00,
            name: "Dona de chocolate a elegir por pieza",
            description: "Dona de chocolate a elegir por pieza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020035100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251986",
            price: 9.00,
            name: "Cuerno por pieza",
            description: "Cuerno por pieza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020020500000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251987",
            price: 49.00,
            name: "Pan infantil paquete con 10 piezas",
            description: "Pan infantil paquete con 10 piezas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020827900000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251988",
            price: 46.00,
            name: "Bisquets paquete con 5 pzas",
            description: "Bisquets paquete con 5 pzas",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020830100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251989",
            price: 9.00,
            name: "Concha de vainilla o chocolate a elegir 1 pieza",
            description: "Concha de vainilla o chocolate a elegir 1 pieza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020010100000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251990",
            price: 19.50,
            name: "Croissant por pieza",
            description: "Croissant por pieza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020020800000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251991",
            price: 19.50,
            name: "Chocolatín por pieza",
            description: "Chocolatín por pieza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020021000000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251992",
            price: 13.50,
            name: "Multi de vainilla pza",
            description: "Multi de vainilla pza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020040300000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251993",
            price: 2.00,
            name: "Bolillo pza",
            description: "Bolillo pza",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00020000300000L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["bakery"]
        },
        {
            sku: "10251994",
            price: 14999.00,
            name: "MacBook Air Apple MGN63LA/A M1 8GB RAM 256GB SSD",
            description: "MacBook Air Apple MGN63LA/A M1 8GB RAM 256GB SSD",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00019425242582l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251995",
            price: 6299.00,
            name: "TV Atvio 50 Pulgadas UHD Smart TV LED ATV-50UHDR",
            description: "TV Atvio 50 Pulgadas UHD Smart TV LED ATV-50UHDR",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00084184810115L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251996",
            price: 7499.00,
            name: "TV Hisense 50 Pulgadas ULED 4K Smart 50U60H",
            description: "TV Hisense 50 Pulgadas ULED 4K Smart 50U60H",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00694214748461l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251997",
            price: 3699.00,
            name: "AirPods 3a Generación Apple Carga Lightning MPNY3AM/A",
            description: "AirPods 3a Generación Apple Carga Lightning MPNY3AM/A",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00019425332403l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251998",
            price: 5499.00,
            name: "Multifuncional Epson Eco Tank L3251 C11CJ67302",
            description: "Multifuncional Epson Eco Tank L3251 C11CJ67302",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00001034395810l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251999",
            price: 10299.00,
            name: "Laptop HP 15-fd0000la Intel Core i3 8 Núcleos 8GB RAM 512GB SSD Azul",
            description: "Laptop HP 15-fd0000la Intel Core i3 8 Núcleos 8GB RAM 512GB SSD Azul",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00019749716428l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251100",
            price: 1499.00,
            name: "Barra de Sonido TCL S332W Negro",
            description: "Barra de Sonido TCL S332W Negro",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00084046409405l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251101",
            price: 7997.00,
            name: "iPhone 11 Apple 64 GB Blanco Telcel",
            description: "iPhone 11 Apple 64 GB Blanco Telcel",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00750622733525l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251102",
            price: 5797.00,
            name: "Smartphone Samsung Galaxy A24 Plata 128 GB Telcel",
            description: "Smartphone Samsung Galaxy A24 Plata 128 GB Telcel",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/gr/images/product-images/img_large/00750622735057L.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        },
        {
            sku: "10251103",
            price: 10497.00,
            name: "iPhone 12 Apple 64 GB Negro Telcel",
            description: "iPhone 12 Apple 64 GB Negro Telcel",
            quantity: 1000,
            image: "https://i5.walmartimages.com.mx/mg/gm/1p/images/product-images/img_large/00750622733945l.jpg?odnHeight=160&odnWidth=160&odnBg=FFFFFF",
            categories: ["computers"]
        }
    ]
);
