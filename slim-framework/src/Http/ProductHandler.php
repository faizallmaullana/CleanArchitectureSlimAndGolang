<?php

declare(strict_types=1);

namespace App\Http;

use App\Application\Product\ProductService;
use App\Domain\Product\ProductNotFound;
use InvalidArgumentException;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Slim\Psr7\Response;

final class ProductHandler
{
    public function __construct(private readonly ProductService $products)
    {
    }

    public function list(ServerRequestInterface $request): ResponseInterface
    {
        return $this->json(['products' => $this->products->list()]);
    }

    public function show(ServerRequestInterface $request, ResponseInterface $response, array $args): ResponseInterface
    {
        try {
            return $this->json($this->products->get($this->productId($args)));
        } catch (ProductNotFound $exception) {
            return $this->error($exception->getMessage(), 404);
        } catch (InvalidArgumentException $exception) {
            return $this->error($exception->getMessage(), 400);
        }
    }

    public function create(ServerRequestInterface $request): ResponseInterface
    {
        try {
            [$name, $price] = $this->payload($request);
            return $this->json($this->products->create($name, $price), 201);
        } catch (InvalidArgumentException $exception) {
            return $this->error($exception->getMessage(), 400);
        }
    }

    public function update(ServerRequestInterface $request, ResponseInterface $response, array $args): ResponseInterface
    {
        try {
            [$name, $price] = $this->payload($request);
            return $this->json($this->products->update($this->productId($args), $name, $price));
        } catch (ProductNotFound $exception) {
            return $this->error($exception->getMessage(), 404);
        } catch (InvalidArgumentException $exception) {
            return $this->error($exception->getMessage(), 400);
        }
    }

    public function delete(ServerRequestInterface $request, ResponseInterface $response, array $args): ResponseInterface
    {
        try {
            $this->products->delete($this->productId($args));
            return new Response(204);
        } catch (ProductNotFound $exception) {
            return $this->error($exception->getMessage(), 404);
        } catch (InvalidArgumentException $exception) {
            return $this->error($exception->getMessage(), 400);
        }
    }

    private function payload(ServerRequestInterface $request): array
    {
        $payload = $request->getParsedBody();

        if (!is_array($payload)) {
            throw new InvalidArgumentException('Request body must be valid JSON.');
        }

        if (!isset($payload['name']) || !is_string($payload['name'])) {
            throw new InvalidArgumentException('Product name is required.');
        }

        if (!isset($payload['price']) || !is_numeric($payload['price'])) {
            throw new InvalidArgumentException('Product price must be numeric.');
        }

        return [$payload['name'], (float) $payload['price']];
    }

    private function productId(array $args): int
    {
        $id = $args['id'] ?? null;

        if (!is_string($id) || !ctype_digit($id) || (int) $id < 1) {
            throw new InvalidArgumentException('Product id must be a positive integer.');
        }

        return (int) $id;
    }

    private function json(mixed $data, int $status = 200): ResponseInterface
    {
        $response = new Response($status);
        $response->getBody()->write((string) json_encode($this->normalize($data)));

        return $response->withHeader('Content-Type', 'application/json');
    }

    private function error(string $message, int $status): ResponseInterface
    {
        return $this->json(['error' => $message], $status);
    }

    private function normalize(mixed $data): mixed
    {
        if (is_object($data) && method_exists($data, 'toArray')) {
            return $data->toArray();
        }

        if (is_array($data)) {
            return array_map(fn (mixed $item): mixed => $this->normalize($item), $data);
        }

        return $data;
    }
}