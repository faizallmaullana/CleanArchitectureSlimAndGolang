<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence;

use App\Domain\Product\Product;
use App\Domain\Product\ProductRepository;
use MongoDB\Collection;

final class MongoProductRepository implements ProductRepository
{
    public function __construct(private readonly Collection $collection)
    {
    }

    public function all(): array
    {
        $documents = $this->collection->find([], [
            'projection' => ['_id' => 0],
            'sort' => ['id' => 1],
        ]);

        $products = [];
        foreach ($documents as $document) {
            $products[] = $this->toProduct($document);
        }

        return $products;
    }

    public function findById(int $id): ?Product
    {
        $document = $this->collection->findOne(
            ['id' => $id],
            ['projection' => ['_id' => 0]]
        );

        return $document === null ? null : $this->toProduct($document);
    }

    public function nextId(): int
    {
        $document = $this->collection->findOne([], [
            'projection' => ['id' => 1],
            'sort' => ['id' => -1],
        ]);

        return $document === null ? 1 : ((int) $document['id']) + 1;
    }

    public function save(Product $product): void
    {
        $this->collection->replaceOne(
            ['id' => $product->id()],
            $product->toArray(),
            ['upsert' => true]
        );
    }

    public function delete(int $id): bool
    {
        return $this->collection->deleteOne(['id' => $id])->getDeletedCount() > 0;
    }

    private function toProduct(mixed $document): Product
    {
        return Product::create(
            (int) $document['id'],
            (string) $document['name'],
            (float) $document['price']
        );
    }
}