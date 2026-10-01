<?php

declare(strict_types=1);

namespace App\Application\Product;

use App\Domain\Product\Product;
use App\Domain\Product\ProductNotFound;
use App\Domain\Product\ProductRepository;

final class ProductService
{
    public function __construct(private readonly ProductRepository $products)
    {
    }

    /** @return list<Product> */
    public function list(): array
    {
        return $this->products->all();
    }

    public function get(int $id): Product
    {
        $product = $this->products->findById($id);

        if ($product === null) {
            throw new ProductNotFound($id);
        }

        return $product;
    }

    public function create(string $name, float $price): Product
    {
        $product = Product::create($this->products->nextId(), $name, $price);
        $this->products->save($product);

        return $product;
    }

    public function update(int $id, string $name, float $price): Product
    {
        $product = $this->get($id);
        $product->update($name, $price);
        $this->products->save($product);

        return $product;
    }

    public function delete(int $id): void
    {
        if (!$this->products->delete($id)) {
            throw new ProductNotFound($id);
        }
    }
}