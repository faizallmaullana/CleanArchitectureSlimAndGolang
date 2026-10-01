<?php

declare(strict_types=1);

require dirname(__DIR__) . '/vendor/autoload.php';

$minimumPrice = isset($argv[1]) ? filter_var($argv[1], FILTER_VALIDATE_FLOAT) : 0.0;

if ($minimumPrice === false || $minimumPrice < 0) {
    fwrite(STDERR, "Minimum price must be a non-negative number.\n");
    exit(1);
}

$client = new MongoDB\Client((string) getenv('MONGODB_URI'));
$databaseName = getenv('MONGODB_DATABASE') ?: 'app_db';
$collection = $client->selectCollection($databaseName, 'products');

$pipeline = [
    ['$match' => ['price' => ['$gte' => $minimumPrice]]],
    ['$project' => ['_id' => 0, 'id' => 1, 'name' => 1, 'price' => 1]],
    ['$sort' => ['price' => -1]],
];

$products = iterator_to_array($collection->aggregate($pipeline), false);
$result = [];

foreach ($products as $product) {
    $result[] = [
        'id' => (int) $product['id'],
        'name' => (string) $product['name'],
        'price' => (float) $product['price'],
    ];
}

echo json_encode([
    'minimum_price' => $minimumPrice,
    'total' => count($result),
    'products' => $result,
], JSON_PRETTY_PRINT) . PHP_EOL;