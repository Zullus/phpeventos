<?php

namespace App\dto;

class RetornaEventoDto
{
    public $id;

    public function __construct($id)
    {
        $this->id = $id;
    }

    public function RetornaEvento()
    {
        $id = $this->id;
    }
}
