<?php

declare(strict_types=1);

namespace App\Domain\Product;

interface ProductRepository
{
    /** @return list<Product> */
    public function all(): array;

    public function findById(int $id): ?Product;

    public function nextId(): int;

    public function save(Product $product): void;

    public function delete(int $id): bool;
}