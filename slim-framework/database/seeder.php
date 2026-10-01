<?php

declare(strict_types=1);

require dirname(__DIR__) . '/vendor/autoload.php';

$client = new MongoDB\Client(getenv('MONGODB_URI'));
$database = getenv('MONGODB_DATABASE') ?: 'app_db';
$collection = $client->$database->products;

$products = [
    ['id' => 1, 'name' => 'Product 1', 'price' => 10.99],
    ['id' => 2, 'name' => 'Product 2', 'price' => 19.99],
    ['id' => 3, 'name' => 'Product 3', 'price' => 5.99],
];

foreach ($products as $product) {
    $collection->replaceOne(
        ['id' => $product['id']],
        $product,
        ['upsert' => true]
    );
}

echo sprintf("Seeded %d products into %s.products\n", count($products), $database);