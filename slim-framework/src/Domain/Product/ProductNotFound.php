<?php

declare(strict_types=1);

namespace App\Domain\Product;

use RuntimeException;

final class ProductNotFound extends RuntimeException
{
    public function __construct(int $id)
    {
        parent::__construct(sprintf('Product with id %d was not found.', $id));
    }
}