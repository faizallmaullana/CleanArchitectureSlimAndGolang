<?php

declare(strict_types=1);

use Slim\Factory\AppFactory;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

require dirname(__DIR__) . '/vendor/autoload.php';

$app = AppFactory::create();

$app->get('/api/products', function (Request $request, Response $response, $args) {
    $body=json_encode([
        'products' => [
            ['id' => 1, 'name' => 'Product 1', 'price' => 10.99],
            ['id' => 2, 'name' => 'Product 2', 'price' => 19.99],
            ['id' => 3, 'name' => 'Product 3', 'price' => 5.99],
        ],
    ]);
    $response->getBody()->write($body);
    return $response->withHeader('Content-Type', 'application/json');
});

$app->run();