<?php

declare(strict_types=1);

use App\Application\Product\ProductService;
use App\Http\ProductHandler;
use App\Infrastructure\Persistence\MongoProductRepository;
use MongoDB\Client;
use Slim\Factory\AppFactory;

require dirname(__DIR__) . '/vendor/autoload.php';

$app = AppFactory::create();
$app->addBodyParsingMiddleware();

$client = new Client((string) getenv('MONGODB_URI'));
$database = getenv('MONGODB_DATABASE') ?: 'app_db';
$repository = new MongoProductRepository($client->selectCollection($database, 'products'));
$handler = new ProductHandler(new ProductService($repository));

$app->get('/api/products', [$handler, 'list']);
$app->get('/api/products/{id}', [$handler, 'show']);
$app->post('/api/products', [$handler, 'create']);
$app->put('/api/products/{id}', [$handler, 'update']);
$app->delete('/api/products/{id}', [$handler, 'delete']);

$app->run();