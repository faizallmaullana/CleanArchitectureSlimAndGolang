<?php

declare(strict_types=1);

namespace App\Domain\Product;

use InvalidArgumentException;

final class Product
{
    private function __construct(
        private readonly int $id,
        private string $name,
        private float $price,
    ) {
        $this->validate($name, $price);
    }

    public static function create(int $id, string $name, float $price): self
    {
        return new self($id, trim($name), $price);
    }

    public function update(string $name, float $price): void
    {
        $name = trim($name);
        $this->validate($name, $price);
        $this->name = $name;
        $this->price = $price;
    }

    public function id(): int
    {
        return $this->id;
    }

    public function toArray(): array
    {
        return [
            'id' => $this->id,
            'name' => $this->name,
            'price' => $this->price,
        ];
    }

    private function validate(string $name, float $price): void
    {
        if ($name === '') {
            throw new InvalidArgumentException('Product name is required.');
        }

        if ($price < 0) {
            throw new InvalidArgumentException('Product price must be zero or greater.');
        }
    }
}